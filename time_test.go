package goods

import (
	"testing"
)

func TestNtpNow(t *testing.T) {
	now, err := NtpNow(``)
	AssertNoError(t, err)
	t.Logf("%v\n", now)
}
