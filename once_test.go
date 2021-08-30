package goods_test

import (
	"testing"

	"github.com/dashessmith/goods"
)

func TestOnce(t *testing.T) {
	var counter int
	goods.Together(func(threadIdx, numThreads int) {
		for i := 0; i < 10000; i++ {
			goods.Once(func() bool {
				return counter <= 0
			}, func() {
				counter++
			})
		}
	})
	goods.AssertEqual(t, counter, 1)
}
