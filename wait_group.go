package util

import (
	"runtime"
	"sync"
)

type WaitGroup struct {
	sync.WaitGroup
}

func (wg *WaitGroup) Go(f func()) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		f()
	}()
}

func (wg *WaitGroup) Together(f func(threadIdx, numThreads int), final func()) {
	numThreads := runtime.NumCPU()
	if final == nil {
		for threadIdx := 0; threadIdx < numThreads; threadIdx++ {
			threadIdx := threadIdx
			wg.Go(func() {
				f(threadIdx, numThreads)
			})
		}
		return
	}
	inner := WaitGroup{}
	inner.Together(f, nil)
	wg.Go(func() {
		inner.Wait()
		final()
	})
}
