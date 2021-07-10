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
	N := 5000000
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		//arr[i] = i
		//arr[i] = i
		arr[i] = rand.Intn(N)
	}
	arr1 := util.CopyInt(arr)
	arr2 := util.CopyInt(arr)
	arr3 := util.CopyInt(arr)
	t.Logf("sort1 start\n")
	d1 := util.Elapse(func() {
		util.MtSort1(arr1, func(i, j int) bool {
			return arr1[i] < arr1[j]
		})
	})
	t.Logf("sort2 start\n")
	d2 := util.Elapse(func() {
		util.MtSort2(arr2, func(i, j int) bool {
			return arr2[i] < arr2[j]
		})
	})
	t.Logf("sort3 start\n")
	d3 := util.Elapse(func() {
		util.MtSort4(arr3, func(i, j int) bool {
			return arr3[i] < arr3[j]
		})
	})
	t.Logf("sort end\n")
	assert.Equal(t, arr1, arr2)
	assert.Equal(t, arr1, arr3)
	assert.True(t, util.MtIsSorted(arr1, func(i, j int) bool { return arr1[i] < arr1[j] }))
	assert.True(t, util.MtIsSorted(arr2, func(i, j int) bool { return arr2[i] < arr2[j] }))
	assert.True(t, util.MtIsSorted(arr3, func(i, j int) bool { return arr3[i] < arr3[j] }))
	t.Logf("d1 = %v, d2 = %v, d3 = %v\n", d1, d2, d3)
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

	N := 1000000
	origin := make([]int, N)
	for i := 0; i < N; i++ {
		origin[i] = rand.Intn(N)
	}
	// origin = []int{7, 9, 3, 9, 0, 7, 2, 4, 0, 8}
	// origin = []int{5, 8, 4, 7, 1, 2, 0, 1, 3, 2}
	arr := util.CopyInt(origin)

	d2 := util.Elapse(func() {
		util.MtSort1(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	})
	if !assert.True(t, sort.IsSorted(util.SortInts(arr))) {
		t.Fatalf("not sorted\norigin %v\nafter %v\n", origin, arr)
	}
	t.Logf("d1 = %v\n", d2)
}

func Test_MtSort4(t *testing.T) {
	rand.Seed(time.Now().Unix())

	N := 100000000
	origin := make([]int, N)
	for i := 0; i < N; i++ {
		origin[i] = rand.Intn(N) / N
		// origin[i] = i
	}
	// origin = []int{7, 9, 3, 9, 0, 7, 2, 4, 0, 8}
	// origin = []int{5, 8, 4, 7, 1, 2, 0, 1, 3, 2}
	arr := util.CopyInt(origin)

	d2 := util.Elapse(func() {
		util.MtSort4(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	})
	if !assert.True(t, sort.IsSorted(util.SortInts(arr))) {
		t.Fatalf("not sorted\norigin %v\nafter %v\n", origin, arr)
	}
	t.Logf("d1 = %v\n", d2)
}

func Test_MtSortOrigin(t *testing.T) {
	rand.Seed(time.Now().Unix())

	N := 100000000
	origin := make([]int, N)
	for i := 0; i < N; i++ {
		origin[i] = rand.Intn(N)
		// origin[i] = i
	}
	// origin = []int{7, 9, 3, 9, 0, 7, 2, 4, 0, 8}
	// origin = []int{5, 8, 4, 7, 1, 2, 0, 1, 3, 2}
	arr := util.CopyInt(origin)

	d2 := util.Elapse(func() {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	})
	if !assert.True(t, sort.IsSorted(util.SortInts(arr))) {
		t.Fatalf("not sorted\norigin %v\nafter %v\n", origin, arr)
	}
	t.Logf("d1 = %v\n", d2)
}

func Test_rand(t *testing.T) {
	t.Logf("rand = %v\n", rand.Intn(100))
}

func Test_isSorted(t *testing.T) {
	N := 50000000
	arr := make([]int, N+1)
	for i := 0; i < N; i++ {
		arr[i] = i
	}
	arr = append(arr, 0)
	r1 := true
	r2 := true
	d1 := util.Elapse(func() {
		r1 = sort.IsSorted(util.SortInts(arr))
	})
	d2 := util.Elapse(func() {
		r2 = util.MtIsSorted(arr, func(i, j int) bool { return arr[i] < arr[j] })
	})
	t.Logf("d1 = %v, d2 = %v\n", d1, d2)
	assert.Equal(t, r1, r2)
}
