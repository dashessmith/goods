package goutil_test

import (
	"testing"

	"github.com/dashessmith/goutil"
)

func Test_WaitGroup(t *testing.T) {
	d := goutil.Elapse(func() {
		wg := goutil.WaitGroup{}
		ch := make(chan int)
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
	t.Logf("elapse %v\n", d)
}
