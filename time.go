package goods

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

// 一日开始时间 2021-5-26 19:20 ==> 2021-5-26 00:00
func TimeFloorDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

// 次日开始时间 2021-5-26 19:20 ==> 2021-5-27 00:00
func TimeCeilDay(t time.Time) time.Time {
	return TimeFloorDay(t).AddDate(0, 0, 1)
}

// 一周开始时间
func TimeFloorWeek(t time.Time) time.Time {
	res := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	switch wd := int(t.Weekday()); wd {
	case 0: // 星期日
		res = res.AddDate(0, 0, -6)
	default: // 周 1 ~ 周 5
		res = res.AddDate(0, 0, -wd+1)
	}
	return res
}

// 下周开始时间
func TimeCeilWeek(t time.Time) time.Time {
	return TimeFloorWeek(t).AddDate(0, 0, 7)
}

// 一月开始时间
func TimeFloorMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
}

// 下月开始时间
func TimeCeilMonth(t time.Time) time.Time {
	return TimeFloorMonth(t).AddDate(0, 1, 0)
}

var nptaddrs = []string{
	"cn.pool.ntp.org",
	"time.windows.com",
	"ntp.aliyun.com",
	"0.beevik-ntp.pool.ntp.org",
}

func NtpNow(addr string) (now time.Time, err error) {
	if len(addr) > 0 {
		return ntp.Time(addr)
	}
	ch := make(chan time.Time, 1)
	for _, addr := range nptaddrs {
		addr := addr
		go func() {
			now, err = ntp.Time(addr)
			if err == nil {
				select {
				case ch <- now:
				default:
				}
			}
		}()
	}
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	select {
	case now = <-ch:
	case <-timer.C:
		err = fmt.Errorf("timeout")
	}
	return
}

func WithDuration(d time.Duration, f func() bool) {
	for t := time.Now(); t.Add(d).After(time.Now()); {
		if f() {
			return
		}
	}
}
