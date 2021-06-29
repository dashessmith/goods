package util_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/dashessmith/util"
	"github.com/stretchr/testify/assert"
)

func Test_MtSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	N := 100000000
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
	}
	arr1 := util.CopyInt(arr)
	arr2 := util.CopyInt(arr)
	d1 := util.Elapse(func() {
		util.MtSort1(arr1, func(i, j int) bool {
			return arr1[i] < arr1[j]
		})
	})
	d2 := util.Elapse(func() {
		util.MtSort2(arr2, func(i, j int) bool {
			return arr2[i] < arr2[j]
		})
	})
	assert.Equal(t, arr1, arr2)
	assert.True(t, sort.IsSorted(util.SortInts(arr1)))
	assert.True(t, sort.IsSorted(util.SortInts(arr1)))
	t.Logf("d1 = %v, d2 = %v\n", d1, d2)
}

func Test_MtSort2(t *testing.T) {
	rand.Seed(time.Now().Unix())
	N := 10
	origin := make([]int, N)
	for i := 0; i < N; i++ {
		origin[i] = rand.Intn(N)
	}
	origin = []int{3, 5, 9, 7, 7, 1, 9, 4, 1, 9}
	arr := util.CopyInt(origin)

	d2 := util.Elapse(func() {
		util.MtSort2(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	})
	if !assert.True(t, sort.IsSorted(util.SortInts(arr))) {
		t.Fatalf("not sorted origin %v, after %v\n", origin, arr)
	}
	t.Logf("d1 = %v\n", d2)
}

func Test_MtSort3(t *testing.T) {
	rand.Seed(time.Now().Unix())

	N := 10
	origin := make([]int, N)
	for i := 0; i < N; i++ {
		origin[i] = rand.Intn(N)
	}
	//origin = []int{7, 9, 3, 9, 0, 7, 2, 4, 0, 8}
	origin = []int{5, 8, 4, 7, 1, 2, 0, 1, 3, 2}
	arr := util.CopyInt(origin)

	d2 := util.Elapse(func() {
		util.MtSort3(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	})
	if !assert.True(t, sort.IsSorted(util.SortInts(arr))) {
		t.Fatalf("not sorted\norigin %v\nafter %v\n", origin, arr)
	}
	t.Logf("d1 = %v\n", d2)
}
