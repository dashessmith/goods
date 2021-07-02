package util_test

import (
	"testing"
	"time"

	"github.com/dashessmith/util"
)

func Test_speed(t *testing.T) {
	go func() {
		util.SpeedSampleF("test1", func() {
			time.Sleep(time.Second)
		})
		for ss := util.NewSpeedSample(`test`); ; {
			ss.Pause()
			time.Sleep(time.Second)
			ss.Resume()
			time.Sleep(time.Second)
			ss.Flush()
		}
	}()
	time.Sleep(30 * time.Second)
}
