package util_test

import (
	"sync"
	"testing"

	"github.com/dashessmith/util"
)

func Test_mtx(t *testing.T) {
	var mtx sync.Mutex
	var a int
	util.Together(func(threadIdx, numThreads int) {
		util.WithMutex(&mtx, func() {
			a++
		})
	})
	t.Logf("cpu = %v\n", a)
}
