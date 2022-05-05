package goods

import (
	"fmt"
	"log"
	"sort"
	"sync"
	"text/tabwriter"
	"time"
)

func speedInit() {
	go func() {
		type S struct {
			samples uint64
			t       float64
			newly   bool
		}
		dict := map[string]*S{}

		lastPrintTime := time.Now()
		for tag := range _speedch {
			s := dict[*tag.tag]
			if s == nil {
				s = &S{}
				dict[*tag.tag] = s
			}
			s.samples++
			s.t += tag.t.Seconds()
			if s.samples == 0 {
				s.samples = 1
				s.t = tag.t.Seconds()
			}
			s.newly = true
			if len(dict) <= 0 || lastPrintTime.Add(3*time.Second).After(time.Now()) {
				continue
			}
			tags := []string{}
			for tag := range dict {
				tags = append(tags, tag)
			}
			sort.Strings(tags)
			writer := tabwriter.NewWriter(log.Default().Writer(), 1, 40, 10, ' ', 0)
			fmt.Fprintf(writer, "--------- speed --------\n")
			fmt.Fprintf(writer, "tag\t\tspeed(sec)\t\tsamples\n")
			for _, tag := range tags {
				s := dict[tag]
				v := s.t / float64(s.samples)
				if s.newly {
					s.newly = false
					fmt.Fprintf(writer, "*%s\t\t%.3f\t\t%v\n", tag, v, s.samples)
				} else {
					fmt.Fprintf(writer, " %s\t\t%.3f\t\t%v\n", tag, v, s.samples)
				}
			}
			writer.Flush()
			lastPrintTime = time.Now()
		}
	}()
}

var (
	_speedch   = make(chan *_Sample, 1024)
	_speedOnce = sync.Once{}
)

type _Sample struct {
	tag *string
	t   time.Duration
}

type SpeedSample struct {
	tag    *string
	beginT time.Time
	endT   time.Time
	pauseT time.Time
}

func NewSpeedSample(tag string) *SpeedSample {
	_speedOnce.Do(speedInit)
	return &SpeedSample{
		tag:    &tag,
		beginT: time.Now(),
	}
}

func (ss *SpeedSample) Pause() {
	ss.pauseT = time.Now()
}

func (ss *SpeedSample) Resume() {
	ss.beginT = ss.beginT.Add(time.Since(ss.pauseT))
}

func (ss *SpeedSample) Flush() {
	ss.endT = time.Now()
	_speedch <- &_Sample{
		tag: ss.tag,
		t:   ss.endT.Sub(ss.beginT),
	}
	ss.beginT = ss.endT
}

func SpeedSampleF(tag string, f func()) {
	_speedOnce.Do(speedInit)
	ss := NewSpeedSample(tag)
	defer ss.Flush()
	if f != nil {
		f()
	}
}
