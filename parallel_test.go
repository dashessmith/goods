package util_test

import (
	"math/rand"
	"testing"

	"github.com/dashessmith/util"
)

func Test_Parall(t *testing.T) {
	t.Logf("test parallel")
	N := 1000000000
	search := 100
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
	}
	d1 := util.Elapse(func() {
		for _, n := range arr {
			if n == search {
				return
			}
		}
	})
	d2 := util.Elapse(func() {
		res := util.AnyOf_P(N, func(i int) bool {
			return arr[i] == search
		})
		t.Logf(" yes ? %v\n ", res)
	})
	t.Logf("d1 = %v, d2 = %v\n", d1, d2)
}
