package goutil_test

import (
	"testing"

	"github.com/dashessmith/goutil"
)

func Test_WaitGroup(t *testing.T) {
	for chsize := 0; chsize <= 10000; chsize += 1000 {
		d := goutil.Elapse(func() {
			wg := goutil.WaitGroup{}
			ch := make(chan int, chsize)
			wg.Together(func(threadIdx, numThreads int) {
				for i := threadIdx; i < 1000000; i += numThreads {
					ch <- i
				}
			}, func() {
				close(ch)
			})
			wg.Together(func(threadIdx, numThreads int) {
				for range ch {

				}
			}, nil)
			wg.Wait()
		})
		t.Logf("chsize %v, elapse %v\n", chsize, d)
	}
}
