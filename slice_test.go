package util_test

import (
	"testing"

	"github.com/dashessmith/util"
)

func Test_ReverseSlice(t *testing.T) {
	var arr []int
	util.ReverseSlice(arr)
	t.Logf("arr = %v\n", arr)
	arr = []int{1}
	util.ReverseSlice(arr)
	t.Logf("arr = %v\n", arr)
	arr = []int{1, 2, 3}
	util.ReverseSlice(arr)
	t.Logf("arr = %v\n", arr)
}
