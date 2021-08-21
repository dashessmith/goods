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

func Test_scalerand(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Logf("%v\n", goods.RandScaleIntWithout(5, 0, 1, 2, 3, 4, 5, 6, 7, 9))
	}
}

func Test_scalerand2(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Logf("%v\n", goods.RandScaleInt(5))
	}
}
