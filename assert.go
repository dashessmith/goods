package goods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertNoError(t *testing.T, err error) {
	if !assert.NoError(t, err) {
		t.Fatal()
	}
}

func AssertTrue(t *testing.T, b bool) {
	if !assert.True(t, b) {
		t.Fatal()
	}
}
