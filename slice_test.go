package goods_test

import (
	"testing"

	"github.com/dashessmith/goods"
)

func Test_ReverseSlice(t *testing.T) {
	var arr []int
	goods.ReverseSlice(arr)
	t.Logf("arr = %v\n", arr)
	arr = []int{1}
	goods.ReverseSlice(arr)
	t.Logf("arr = %v\n", arr)
	arr = []int{1, 2, 3}
	goods.ReverseSlice(arr)
	t.Logf("arr = %v\n", arr)
}
