package util_test

import (
	"testing"

	"github.com/dashessmith/util"
)

func Test_stringsjoin(t *testing.T) {
	s := util.JoinInts([]int{1, 2, 3}, "-")
	t.Logf("%v\n", s)
}
