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
	if atomic.AddInt64(&wg.goCount, 1) < int64(runtime.NumCPU()) {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				atomic.AddInt64(&wg.goCount, -1)
			}()
			f()
		}()
	} else {
		atomic.AddInt64(&wg.goCount, -1)
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
