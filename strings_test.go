package goods_test

import (
	"testing"

	"github.com/dashessmith/goods"
)

func Test_stringsjoin(t *testing.T) {
	s := goods.JoinInts([]int{1, 2, 3}, "-")
	t.Logf("%v\n", s)
}
