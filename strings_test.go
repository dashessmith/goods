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
