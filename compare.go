package goods

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(ints []int) (ret int) {
	for _, n := range ints {
		if n > ret {
			ret = n
		}
	}
	return
}
