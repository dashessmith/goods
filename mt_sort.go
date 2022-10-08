package goods

import (
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
)

const (
	MTSORT_THREADLIMIT_FOR_INTS = (1 << 10)
)

type MtSortOption struct {
	ThreadLimit       int
	SkipCheckIsSorted bool
}

// MtSort sort in multi threads
func MtSort(x interface{}, less func(i, j int) bool, opt ...*MtSortOption) {
	var o *MtSortOption
	if len(opt) > 0 {
		o = opt[0]
	} else {
		o = &MtSortOption{
			SkipCheckIsSorted: true,
		}
	}
	if o.ThreadLimit < 0 {
		o.ThreadLimit = 0
	}
	MtSort4(x, less, o)
}

func MtIsSorted(x interface{}, less func(i, j int) bool) bool {
	N := reflect.ValueOf(x).Len()
	var res uint32 = 1
	edgeIdx := []int{}
	edgeMask := map[int]bool{}
	edgeMtx := sync.Mutex{}
	Together(func(threadIdx, numThreads int) {
		tN := N / numThreads
		start := threadIdx * tN
		end := (threadIdx + 1) * tN
		if end > N {
			end = N
		}
		WithMutex(&edgeMtx, func() {
			if start < N && !edgeMask[start] {
				edgeMask[start] = true
				edgeIdx = append(edgeIdx, start)
			}
			if end-1 < N && !edgeMask[end-1] {
				edgeMask[end-1] = true
				edgeIdx = append(edgeIdx, end-1)
			}
		})
		for i := start + 1; atomic.LoadUint32(&res) > 0 && i < end; i++ {
			if less(i, i-1) {
				atomic.StoreUint32(&res, 0)
				break
			}
		}
	})
	if res == 0 {
		return false
	}
	sort.Slice(edgeIdx, func(i, j int) bool {
		return edgeIdx[i] < edgeIdx[j]
	})
	for i := 1; i < len(edgeIdx); i++ {
		if less(i, i-1) {
			return false
		}
	}
	return true
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

// MtSort4 q sort mt mid pivot, first rand pivot
func MtSort4(x interface{}, less func(i, j int) bool, opt *MtSortOption) {
	if !opt.SkipCheckIsSorted && MtIsSorted(x, less) {
		return
	}
	swap := reflect.Swapper(x)
	N := reflect.ValueOf(x).Len()
	if N <= 1 {
		return
	}
	wg := WaitGroup{}
	defer wg.Wait()
	first := true
	var impl func(left, right int)
	impl = func(left, right int) {
		if left >= right {
			return
		}
		if left+1 == right {
			if less(right, left) {
				swap(right, left)
			}
			return
		}
		var pivot int
		if first {
			pivot = left + rand.Intn(right-left)
			first = false
		} else {
			pivot = (left + right) / 2
		}
		swap(pivot, right)
		j := left
		for i := left; i < right; i++ {
			if less(i, right) {
				swap(i, j)
				j++
			}
		}
		swap(j, right)
		pivot = j

		if pivot-left > right-pivot-1 {
			if left < pivot {
				if left+opt.ThreadLimit < pivot {
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
				if pivot+1+opt.ThreadLimit < right {
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
	impl(0, N-1)
}
