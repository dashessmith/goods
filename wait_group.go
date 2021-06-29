package util

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type WaitGroup struct {
	sync.WaitGroup
	goCount int64
}

func (wg *WaitGroup) GoOrCall(f func()) {
	defer atomic.AddInt64(&wg.goCount, -1)
	if atomic.AddInt64(&wg.goCount, 1) < int64(runtime.NumCPU()) {
		wg.Go(f)
	} else {
		f()
	}
}

func (wg *WaitGroup) Go(f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}

func (wg *WaitGroup) Together(f func(threadIdx, numThreads int)) {
	numThreads := runtime.NumCPU()
	for threadIdx := 0; threadIdx < numThreads; threadIdx++ {
		threadIdx := threadIdx
		wg.Go(func() {
			f(threadIdx, numThreads)
		})
	}
}

func (wg *WaitGroup) TogetherF(f func(threadIdx, numThreads int), final func()) {
	if final == nil {
		wg.Together(f)
		return
	}
	inner := WaitGroup{}
	inner.Together(f)
	wg.Go(func() {
		inner.Wait()
		final()
	})
}
