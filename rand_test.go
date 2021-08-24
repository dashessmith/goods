package goods

import "testing"

func TestRandScaleStr(t *testing.T) {
	t.Logf("%v\n", RandScaleStr(2))
	t.Logf("%v\n", RandScaleStr(3))
	t.Logf("%v\n", RandScaleStr(4))
	t.Logf("%v\n", RandScaleStr(5))
}
