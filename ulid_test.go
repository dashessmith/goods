package util_test

import (
	"testing"

	"github.com/dashessmith/util"
)

func Test_ulid(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Logf("%v\n", util.ULID())
	}
}

func Test_ulidint64(t *testing.T) {
	for i := 0; i < 100; i++ {
		t.Logf("%v\n", util.ULIDInt64())
	}
}
