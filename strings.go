package goods

import "strconv"

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
