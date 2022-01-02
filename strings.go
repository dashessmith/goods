package goods

import (
	"regexp"
	"strconv"
	"strings"
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

func ReplaceAll(s string, new string, olds ...string) string {
	for _, old := range olds {
		s = strings.ReplaceAll(s, old, new)
	}
	return s
}

func TrimIf(s string, precond func(r rune) bool) string {
	ret := []rune(s)
	return string(RemoveIf(ret, func(idx int) bool {
		return precond(ret[idx])
	}).([]rune))
}

func Grep(src string, pattern string) (ret string, err error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	matches := r.FindStringSubmatch(src)
	if len(matches) >= 2 {
		ret = matches[1]
	}
	return
}

func GrepInt(src string, pattern string) (ret int, err error) {
	str, err := Grep(src, pattern)
	if err != nil {
		return
	}
	return strconv.Atoi(str)
}
