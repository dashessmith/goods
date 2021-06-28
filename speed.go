package util

import (
	"fmt"
	"log"
	"sort"
	"text/tabwriter"
	"time"
)

func init() {
	go func() {
		type S struct {
			samples int
			t       time.Time
		}
		dict := map[string]*S{}

		lastPrintTime := time.Now()
		for tag := range _speedch {
			s := dict[*tag.tag]
			if s == nil {
				s = &S{
					samples: 1,
					t:       tag.t,
				}
				dict[*tag.tag] = s
			} else {
				s.samples++
			}
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
			fmt.Fprintf(writer, "tag\t\tspeed(/sec)\n")
			for _, tag := range tags {
				s := dict[tag]
				fmt.Fprintf(writer, "%s\t\t%.2f\n", tag, float64(s.samples)/time.Since(s.t).Seconds())
				delete(dict, tag)
			}
			writer.Flush()
			lastPrintTime = time.Now()
		}
	}()
}

var _speedch = make(chan *_Sample, 1024)

type _Sample struct {
	tag *string
	t   time.Time
}

func Speed(tag string) {
	_speedch <- &_Sample{
		tag: &tag,
		t:   time.Now(),
	}
}
