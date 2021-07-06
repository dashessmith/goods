package util_test

import (
	"testing"

	"github.com/dashessmith/util"
)

func Test_encrypt(t *testing.T) {
	str := "hello world"
	key := [32]byte{0}
	encrypted := util.Encrypt([]byte(str), key[:])
	t.Logf("%s\n", encrypted)
	decrypted := util.Decrypt(encrypted, key[:])
	t.Logf("%s\n", decrypted)
}
