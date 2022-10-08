package goods

import (
	"testing"
)

func TestIsCnMobile(t *testing.T) {
	AssertTrue(t, IsCnMobile("13000000000"))
	AssertTrue(t, !IsCnMobile("12000000000"))
	AssertTrue(t, !IsCnMobile("11000000000"))
	AssertTrue(t, !IsCnMobile("10000000000"))
	AssertTrue(t, !IsCnMobile("1300000000"))
	AssertTrue(t, IsCnMobile("+8613000000000"))
	AssertTrue(t, !IsCnMobile("+861300000000"))
}
