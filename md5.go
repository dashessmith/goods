package util

import (
	"crypto/md5"
	"encoding/hex"
)

func GenMD5(input string) string {
	var hash = md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
