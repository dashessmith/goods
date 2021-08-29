package goods_test

import (
	"math/rand"
	"testing"

	"github.com/dashessmith/goods"
)

func Test_Parall(t *testing.T) {
	t.Logf("test parallel")
	N := 10000000
	search := 100
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
	}
	d1 := goods.Elapse(func() {
		for _, n := range arr {
			if n == search {
				return
			}
		}
	})
	d2 := goods.Elapse(func() {
		res := goods.AnyOf_P(N, func(i int) bool {
			return arr[i] == search
		})
		t.Logf(" yes ? %v\n ", res)
	})
	t.Logf("d1 = %v, d2 = %v\n", d1, d2)
}
