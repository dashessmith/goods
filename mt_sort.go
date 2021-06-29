package util

import (
	"reflect"
	"runtime"
	"sync/atomic"
)

// MtSort sort in multi threads
func MtSort(x interface{}, less func(i, j int) bool) {
	MtSort1(x, less)
}

// MtSort sort in multi threads
func MtSort1(x interface{}, less func(i, j int) bool) {
	threadLimit := (1 << 10)
	swap := reflect.Swapper(x)
	wg := WaitGroup{}
	defer wg.Wait()
	var impl func(left, right int)
	impl = func(left, right int) {
		pivot := left
		for l, r := left, right; l < r; {
			for ; pivot < r && !less(r, pivot); r-- {
			}
			if pivot >= r {
				break
			}
			swap(pivot, r)
			pivot = r
			for ; l < pivot && !less(pivot, l); l++ {
			}
			if l >= pivot {
				break
			}
			swap(pivot, l)
			pivot = l
		}
		if pivot-left > right-pivot-1 {
			if left < pivot {
				if left+threadLimit < pivot {
					wg.GoOrCall(func() {
						impl(left, pivot)
					})
				} else {
					impl(left, pivot)
				}
			}
			if pivot+1 < right {
				impl(pivot+1, right)
			}
		} else {
			if pivot+1 < right {
				if pivot+1+threadLimit < right {
					wg.GoOrCall(func() {
						impl(pivot+1, right)
					})
				} else {
					impl(pivot+1, right)
				}
			}
			if left < pivot {
				impl(left, pivot)
			}
		}
	}
	impl(0, reflect.ValueOf(x).Len()-1)
}

// MtSort2 sort in multi threads
func MtSort2(x interface{}, less func(i, j int) bool) {
	swap := reflect.Swapper(x)
	wg := WaitGroup{}
	defer wg.Wait()

	ch := make(chan []int, runtime.NumCPU())

	totalN := reflect.ValueOf(x).Len()
	if totalN <= 1 {
		return
	}
	var N int64

	incre := func(n int64) {
		if atomic.AddInt64(&N, n) >= int64(totalN) {
			SafeClose(ch)
		}
	}

	var partition func(left, right int)
	partition = func(left, right int) {
		pivot := left
		for l, r := left, right; l < r; {
			for ; pivot < r && !less(r, pivot); r-- {
			}
			if pivot >= r {
				break
			}
			swap(pivot, r)
			pivot = r
			for ; l < pivot && !less(pivot, l); l++ {
			}
			if l >= pivot {
				break
			}
			swap(pivot, l)
			pivot = l
		}
		if pivot-left > right-pivot-1 {
			if left < pivot {
				select {
				case ch <- []int{left, pivot}:
				default:
					partition(left, pivot)
				}
			} else {
				incre(1)
			}
			if pivot+1 < right {
				select {
				case ch <- []int{pivot + 1, right}:
				default:
					partition(pivot+1, right)
				}
			} else if pivot+1 == right {
				incre(1)
			}
		} else {
			if pivot+1 < right {
				select {
				case ch <- []int{pivot + 1, right}:
				default:
					partition(pivot+1, right)
				}
			} else if pivot+1 == right {
				incre(1)
			}

			if left < pivot {
				select {
				case ch <- []int{left, pivot}:
				default:
					partition(left, pivot)
				}
			} else {
				incre(1)
			}
		}
	}
	wg.Together(func(threadIdx, numThreads int) {
		for lr := range ch {
			partition(lr[0], lr[1])
		}
	})

	ch <- []int{0, totalN - 1}
}

// merge sort
func MtSort3(x interface{}, less func(i, j int) bool) {
	swap := reflect.Swapper(x)

	N := reflect.ValueOf(x).Len()

	nextGap := func(gap int) int {
		if gap <= 1 {
			return 0
		}
		return gap/2 + gap%2
	}
	var impl func(l, r int)
	impl = func(l, r int) {
		if l+1 >= r {
			return
		}
		if l+2 == r {
			if less(r-1, l) {
				swap(r-1, l)
			}
			return
		}
		mid := (l + r) / 2
		if r-l >= 2048 {
			wg := WaitGroup{}
			wg.Go(func() {
				impl(l, mid)
			})
			impl(mid, r)
			wg.Wait()
		} else {
			impl(l, mid)
			impl(mid, r)
		}

		for gap := nextGap(r - l); gap > 0; gap = nextGap(gap) {
			for i := l; i+gap < r; i++ {
				if less(i+gap, i) {
					swap(i+gap, i)
				}
			}
		}
	}
	impl(0, N)
}
