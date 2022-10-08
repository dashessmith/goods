package std_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/dashessmith/goods/std"
)

func TestList(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	l := std.List[int]{}

	// {
	// 	l.PushBack(10)
	// 	l.PushBack(1)
	// 	l.PushBack(0)
	// 	l.PushBack(3)
	// 	l.PushBack(4)
	// }

	randnums := map[int]int{}
	originList := []int{}
	pushl := func(x int) {
		l.PushBack(x)
		originList = append(originList, x)
		randnums[x]++
	}

	// for _, n := range []int{76, 2275, 3727, 4176, 1171, 5885, 3298, 5773, 7476, 2685} {
	// for _, n := range []int{9370, 295, 9520, 940, 1429, 2114, 7459, 9935, 8588, 1729} {
	// for _, n := range []int{0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0} {
	// 	pushl(n)
	// }

	// pushl(1)
	// pushl(0)
	// pushl(3)

	for i := 0; i < 10000000; i++ {
		// r := rand.Intn(2)
		// r := rand.Intn(10000000)
		// r := i
		r := -i
		pushl(r)
	}

	// for _, x := range []int{0, 5, 3, 3, 2, 2, 2, 5, 0} {
	// for _, x := range []int{81, 88, 223, 306, 318, 387, 411, 406, 440, 462, 655, 704, 1026, 1039, 1144, 1313, 1370, 1378, 1613, 1644, 1690, 1742, 1885, 1888, 1925, 1935, 1943, 2162, 2212, 2213, 2227, 2268, 2287, 2442, 2521, 2539, 2683, 2886, 3455, 2901, 3127, 3506, 3656, 4129, 4151, 4203, 4411, 4552, 4740, 4874, 4915, 4971, 5194, 5294, 5330, 5403, 5487, 5569, 5671, 5687, 5870, 6162, 6209, 6228, 6253, 6484, 6557, 6665, 6868, 6970, 6957, 7075, 7222, 7241, 7476, 7520, 7666, 7880, 8021, 8120, 8275, 8326, 8522, 8687, 8723, 8880, 8916, 8960, 8977, 9082, 9145, 9168, 9195, 9216, 9255, 9288, 9329, 9331, 9805, 9905} {
	// 	pushl(x)
	// }

	// for i := 0; i < 3; i++ {
	// 	l.PushBack(i)
	// }
	beginTime := time.Now()
	l.Sort(func(a, b *int) bool {
		return *a < *b
	})
	t.Logf("sort time %v", time.Since(beginTime))

	iterNums := map[int]int{}

	for itr := l.Begin(); !itr.Equal(l.End()); itr = itr.Next() {
		iterNums[*itr.Value()]++
		// t.Logf("%v", *itr.Value())
	}

	for k, v := range iterNums {
		if randnums[k] != v {
			t.Fatalf("not equal")
		}
	}

	for k, v := range randnums {
		if iterNums[k] != v {
			t.Fatalf("not equal")
		}
	}

	var lastValue *int
	for itr := l.Begin(); !itr.Equal(l.End()); itr = itr.Next() {
		if lastValue == nil {
			lastValue = itr.Value()
		} else if *lastValue > *itr.Value() {
			// fmt.Printf("origin list ")
			// for _, n := range originList {
			// 	fmt.Printf("%v,", n)
			// }
			// fmt.Printf("\n")
			t.Fatalf("not ordered %v > %v", *lastValue, *itr.Value())
		}
		*lastValue = *itr.Value()
	}
}
