package goods

import "testing"

func TestRename(t *testing.T) {
	x := Rename("C:/Users/z/test.txt", "c:/temp/test.txt")
	AssertNoError(t, x)
}

func TestCopy(t *testing.T) {
	x := Copy("C:/temp/test.txt", "c:/temp/test2.txt")
	AssertNoError(t, x)
}
