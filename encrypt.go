package goods

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func IrreversibleEncryptStr(input, key string) string {
	md5e := md5.Sum([]byte(input + key))
	return hex.EncodeToString(md5e[:])
}

func fixkey(key []byte) (output []byte) {
	output = make([]byte, 32)
	copy(output, key)
	return
}

// only take front 32 bytes of key, or '\0' appended
func Encrypt(plaintext []byte, key []byte) (encryptedString []byte, err error) {
	key = fixkey(key)

	// Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	// https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	// Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	// Encrypt the data using aesGCM.Seal
	// Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

// only take front 32 bytes of key, or '\0' appended
func Decrypt(enc, key []byte) (decryptedString []byte, err error) {
	key = fixkey(key)
	// Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	// Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	// Get the nonce size
	nonceSize := aesGCM.NonceSize()

	// Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return
	}

	return plaintext, nil
}

func EncryptStr(src, key string) (string, error) {
	e, err := Encrypt([]byte(src), []byte(key))
	if err != nil {
		return ``, err
	}
	return hex.EncodeToString(e), nil
}

func DecryptStr(src, key string) (string, error) {
	srcbytes, err := hex.DecodeString(src)
	if err != nil {
		return ``, err
	}
	d, err := Decrypt(srcbytes, []byte(key))
	if err != nil {
		return ``, err
	}
	return string(d), nil
}
