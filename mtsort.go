package util

import (
	"reflect"
)

// MtSort sort in multi threads
func MtSort(x interface{}, less func(i, j int) bool) {
	threadLimit := (1 << 12)
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
		if left < pivot {
			if left+threadLimit < pivot {
				wg.Go(func() {
					impl(left, pivot)
				})
			} else {
				impl(left, pivot)
			}
		}
		if pivot+1 < right {
			impl(pivot+1, right)
		}
	}
	impl(0, reflect.ValueOf(x).Len()-1)
}
