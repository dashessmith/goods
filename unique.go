package util

func UniqueInt64(srcLength int, get func(index int) int64) (res []int64) {
	mask := map[int64]bool{}
	for i := 0; i < srcLength; i++ {
		if n := get(i); mask[n] {
			continue
		} else {
			mask[n] = true
			res = append(res, n)
		}
	}
	return
}
