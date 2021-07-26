package goods_test

import (
	"sync"
	"testing"

	"github.com/dashessmith/goods"
)

func Test_mtx(t *testing.T) {
	var mtx sync.Mutex
	var a int
	goods.Together(func(threadIdx, numThreads int) {
		goods.WithMutex(&mtx, func() {
			a++
		})
	})
	t.Logf("cpu = %v\n", a)
}
