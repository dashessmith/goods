package goods_test

import (
	"testing"

	"github.com/dashessmith/goods"
)

func Test_ulid(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Logf("%v\n", goods.ULID())
	}
}

func Test_ulidint64(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Logf("%v\n", goods.ULIDInt64())
	}
}
