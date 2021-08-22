package goods

func UniqueInt64f(srcLength int, get func(index int) int64) (res []int64) {
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

func UniqueInt64(src []int64) (res []int64) {
	return UniqueInt64f(len(src), func(index int) int64 { return src[index] })
}

func UniqueIntf(srcLength int, get func(index int) int, mask map[int]bool) (res []int) {
	if mask == nil {
		mask = map[int]bool{}
	}
	for i := 0; i < srcLength; i++ {
		if n := get(i); !mask[n] {
			mask[n] = true
			res = append(res, n)
		}
	}
	return
}

func UniqueInt(src []int) (res []int) {
	return UniqueIntf(len(src), func(index int) int { return src[index] }, nil)
}

func UniqueStrings(strs []string) (res []string) {
	mask := map[string]bool{}
	for _, str := range strs {
		if mask[str] {
			continue
		}
		mask[str] = true
		res = append(res, str)
	}
	return
}
