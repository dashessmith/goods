package goods

import "testing"

func TestAllIntsEqual(t *testing.T) {
	ints := []int{1, 1, 1}
	AssertTrue(t, AllIntsEqual(ints...))
	ints = []int{}
	AssertTrue(t, AllIntsEqual(ints...))
	ints = []int{1, 0}
	AssertFalse(t, AllIntsEqual(ints...))
	ints = []int{1, 0}
	AssertEqual(t, SumInts(ints...), 1)
	ints = []int{1, 0, -1}
	AssertEqual(t, SumInts(ints...), 0)
	ints = []int{}
	AssertEqual(t, SumInts(ints...), 0)
}
