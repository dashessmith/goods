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
