package goods

import "testing"

func TestRename(t *testing.T) {
	x := Rename("C:/Users/z/test.txt", "c:/temp/test.txt")
	AssertNoError(t, x)
}
