package goods

import "reflect"

func ReverseSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	N := rv.Len()
	for i, j := 0, N-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func RemoveIf(slice interface{}, precond func(idx int) bool, removecount ...*int) (ret interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()

	keep := 0
	idx := 0
	for ; idx < length; idx++ {
		if precond(idx) {
			continue
		}
		if idx != keep {
			swap(idx, keep)
		}
		keep++
	}
	if len(removecount) > 0 {
		*removecount[0] = length - keep
	}
	return rv.Slice(0, keep).Interface()
}

func Any(slice interface{}, precond func(idx int) bool) bool {
	rv := reflect.ValueOf(slice)
	length := rv.Len()
	for idx := 0; idx < length; idx++ {
		if precond(idx) {
			return true
		}
	}
	return false
}
