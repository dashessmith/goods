package util_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/dashessmith/util"
	"github.com/magiconair/properties/assert"
)

func Test_MtSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	N := 10000000
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
	}
	arr1 := util.CopyInt(arr)
	arr2 := util.CopyInt(arr)

	d1 := util.Elapse(func() {
		sort.Slice(arr1, func(i, j int) bool {
			return arr1[i] < arr1[j]
		})
	})
	d2 := util.Elapse(func() {
		util.MtSort(arr2, func(i, j int) bool {
			return arr2[i] < arr2[j]
		})
	})
	assert.Equal(t, arr1, arr2)
	t.Logf("d1 = %v, d2 = %v\n", d1, d2)
}
