package goods

import "testing"

func foo() string {
	return MyCaller()
}

func TestMyCaller(t *testing.T) {
	t.Logf("%v\n", foo())
}
