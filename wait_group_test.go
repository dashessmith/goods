package goods_test

import (
	"testing"

	"github.com/dashessmith/goods"
)

func Test_WaitGroup(t *testing.T) {
	for chsize := 0; chsize <= 10000; chsize += 1000 {
		d := goods.Elapse(func() {
			wg := goods.WaitGroup{}
			ch := make(chan int, chsize)
			wg.TogetherF(func(threadIdx, numThreads int) {
				for i := threadIdx; i < 1000000; i += numThreads {
					ch <- i
				}
			}, func() {
				close(ch)
			})
			wg.Together(func(threadIdx, numThreads int) {
				for range ch {
				}
			})
			wg.Wait()
		})
		t.Logf("chsize %v, elapse %v\n", chsize, d)
	}
}
