package goods_test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/dashessmith/goods"
)

func Test_MtSort(t *testing.T) {
	rand.Seed(time.Now().Unix())
	N := 100000000
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		// arr[i] = rand.Intn(N)
		arr[i] = i
	}

	tsort := func(tag string, sortf func(interface{}, func(int, int) bool, ...*goods.MtSortOption)) {
		arrcp := goods.CopyInt(arr)
		t.Logf("%v start\n", tag)
		d3 := goods.Elapse(func() {
			sortf(arrcp, func(i, j int) bool {
				return arrcp[i] < arrcp[j]
			})
		})
		goods.AssertTrue(t, goods.MtIsSorted(arrcp, func(i, j int) bool { return arrcp[i] < arrcp[j] }))
		t.Logf("%v = %v\n", tag, d3)
	}
	// tsort(`mt sort 1`, goods.MtSort1)
	// tsort(`mt sort 2`, goods.MtSort2)
	tsort(`mt sort 4`, goods.MtSort)

	t.Logf("sort end\n")
}

func Test_MtSort2(t *testing.T) {
	rand.Seed(time.Now().Unix())
	N := 10
	origin := make([]int, N)
	for i := 0; i < N; i++ {
		origin[i] = rand.Intn(N)
	}
	origin = []int{3, 5, 9, 7, 7, 1, 9, 4, 1, 9}
	arr := goods.CopyInt(origin)

	d2 := goods.Elapse(func() {
		goods.MtSort2(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	})
	goods.AssertTrue(t, sort.IsSorted(goods.SortInts(arr)))
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
	arr := goods.CopyInt(origin)

	d2 := goods.Elapse(func() {
		goods.MtSort1(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	})
	goods.AssertTrue(t, sort.IsSorted(goods.SortInts(arr)))
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
	arr := goods.CopyInt(origin)

	d2 := goods.Elapse(func() {
		goods.MtSort4(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		}, &goods.MtSortOption{
			ThreadLimit: goods.MTSORT_THREADLIMIT_FOR_INTS,
		})
	})
	goods.AssertTrue(t, sort.IsSorted(goods.SortInts(arr)))
	t.Logf("d2 = %v\n", d2)
}

func Test_MtSortOrigin(t *testing.T) {
	rand.Seed(time.Now().Unix())

	N := 1000000
	origin := make([]int, N)
	for i := 0; i < N; i++ {
		origin[i] = rand.Intn(N)
		// origin[i] = i
	}
	// origin = []int{7, 9, 3, 9, 0, 7, 2, 4, 0, 8}
	// origin = []int{5, 8, 4, 7, 1, 2, 0, 1, 3, 2}
	arr := goods.CopyInt(origin)

	d2 := goods.Elapse(func() {
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] < arr[j]
		})
	})
	goods.AssertTrue(t, sort.IsSorted(goods.SortInts(arr)))
	t.Logf("d1 = %v\n", d2)
}

func Test_rand(t *testing.T) {
	t.Logf("rand = %v\n", rand.Intn(100))
}

func Test_isSorted(t *testing.T) {
	N := 50000000
	arr := make([]int, N+1)
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(N)
	}
	arr = append(arr, 0)
	r1 := true
	r2 := true
	d1 := goods.Elapse(func() {
		r1 = sort.IsSorted(goods.SortInts(arr))
	})
	d2 := goods.Elapse(func() {
		r2 = goods.MtIsSorted(arr, func(i, j int) bool { return arr[i] < arr[j] })
	})
	t.Logf("d1 = %v, d2 = %v\n", d1, d2)
	goods.AssertEqual(t, r1, r2)
}
