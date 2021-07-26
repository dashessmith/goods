package goods_test

import (
	"testing"

	"github.com/dashessmith/goods"
)

func Test_rt(t *testing.T) {
	t.Logf("%v\n, %v\n, %v\n, %v\n, %v\n", goods.BinName, goods.BinExt, goods.BinNameExt, goods.BinDir, goods.BinPath)
}
