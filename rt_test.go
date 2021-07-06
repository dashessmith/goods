package util_test

import (
	"testing"
	"github.com/dashessmith/util"
)

func Test_rt(t *testing.T) {
	t.Logf("%v\n, %v\n, %v\n, %v\n, %v\n", util.BinName, util.BinExt, util.BinNameExt, util.BinDir, util.BinPath)
}
