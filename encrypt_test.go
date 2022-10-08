package goods_test

import (
	"testing"

	"github.com/dashessmith/goods"
)

func Test_encrypt(t *testing.T) {
	str := "hello world"
	key := [32]byte{0}
	encrypted, err := goods.Encrypt([]byte(str), key[:])
	goods.AssertNoError(t, err)
	t.Logf("%s\n", encrypted)
	decrypted, err := goods.Decrypt(encrypted, key[:])
	goods.AssertNoError(t, err)
	t.Logf("%s\n", decrypted)
}

func Test_encrypt2(t *testing.T) {
	str := "hello world"
	key := `123`
	encrypted, err := goods.EncryptStr(str, key)
	goods.AssertNoError(t, err)
	t.Logf("%s\n", encrypted)
	decrypted, err := goods.DecryptStr(encrypted, key)
	goods.AssertNoError(t, err)
	t.Logf("%s\n", decrypted)
	goods.AssertEqual(t, str, decrypted)
}

func Test_encryptstr(t *testing.T) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 32; j++ {
			str := goods.RandScaleStr(i + 1)
			key := goods.RandScaleStr(j)
			encrypted, err := goods.EncryptStr(str, key)
			goods.AssertNoError(t, err)
			decrypted, err := goods.DecryptStr(encrypted, key)
			goods.AssertNoError(t, err)
			goods.AssertEqual(t, str, decrypted)
			goods.AssertTrue(t, str != encrypted)
		}
	}
}

func TestIrreversibleEncryptStr(t *testing.T) {
	x := goods.IrreversibleEncryptStr("123", "1")
	goods.AssertEqual(t, x, `6c14da109e294d1e8155be8aa4b1ce8e`)
}
