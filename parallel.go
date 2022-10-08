package goods

func AnyOf_P(N int, f func(i int) bool) (yes bool) {
	wg := WaitGroup{}
	defer wg.Wait()
	wg.Together(func(threadIdx, numThreads int) {
		for i := threadIdx; !yes && i < N; i += numThreads {
			if f(i) {
				yes = true
			}
		}
	})
	return
}

func NoneOf_P(N int, f func(i int) bool) (yes bool) {
	wg := WaitGroup{}
	defer wg.Wait()
	yes = true
	wg.Together(func(threadIdx, numThreads int) {
		for i := threadIdx; yes && i < N; i += numThreads {
			if f(i) {
				yes = false
			}
		}
	})
	return
}

func Foreach_P(N int, f func(i int)) {
	wg := WaitGroup{}
	defer wg.Wait()
	wg.Together(func(threadIdx, numThreads int) {
		for i := threadIdx; i < N; i += numThreads {
			f(i)
		}
	})
}
