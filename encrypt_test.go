package goods_test

import (
	"testing"

	"github.com/dashessmith/goods"
)

func Test_encrypt(t *testing.T) {
	str := "hello world"
	key := [32]byte{0}
	encrypted := goods.Encrypt([]byte(str), key[:])
	t.Logf("%s\n", encrypted)
	decrypted := goods.Decrypt(encrypted, key[:])
	t.Logf("%s\n", decrypted)
}
