package goods_test

import (
	"testing"

	"github.com/dashessmith/goods"
)

func Test_stringsjoin(t *testing.T) {
	s := goods.JoinInts([]int{1, 2, 3}, "-")
	t.Logf("%v\n", s)
}

func TestTrimAllSpace(t *testing.T) {
	goods.AssertEqual(t, goods.TrimAllSpace("1 2 3"), "123")
}

func TestTrimIf(t *testing.T) {
	nums := `1 2 3  5   6 7 \t 890`
	nums = goods.TrimIf(nums, func(r rune) bool {
		return r < '0' || r > '9'
	})
	t.Logf("%v\n", nums)
}
