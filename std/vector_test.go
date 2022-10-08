package std

import "testing"

func TestVector(t *testing.T) {
	v := Vector[int]{}
	v.InsertBefore(0, 1)
	v.InsertBefore(0, 2)
	v.InsertBefore(1, 3)
	t.Logf("v = %v", v)
	v.EraseAt(1)
	t.Logf("v = %v", v)
	v.Clear()
	t.Logf("v = %v", v)
}
