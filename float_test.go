package goods

import "testing"

func TestFixedFloat64(t *testing.T) {
	AssertEqual(t, FixedFloat64(1.2345678, 0), float64(1)) // 1
	AssertEqual(t, FixedFloat64(1.2345678, 1), 1.2)        // 1.2
	AssertEqual(t, FixedFloat64(1.2345678, 2), 1.23)       // 1.23
	AssertEqual(t, FixedFloat64(1.2345678, 3), 1.235)      // 1.235 (rounded up)
}
