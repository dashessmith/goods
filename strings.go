package goods

import (
	"strconv"
	"unicode"
)

func JoinInts(n []int, sep string) (res string) {
	for x := range n {
		if len(res) > 0 {
			res += sep + strconv.Itoa(x)
		} else {
			res = strconv.Itoa(x)
		}
	}
	return
}

func TrimAllSpace(x string) (ret string) {
	retrune := []rune{}
	for _, r := range x {
		if !unicode.IsSpace(r) {
			retrune = append(retrune, r)
		}
	}
	return string(retrune)
}
