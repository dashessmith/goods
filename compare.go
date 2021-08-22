package goods

import "math"

func MinInt(ints ...int) (ret int) {
	ret = math.MaxInt
	for _, n := range ints {
		if n < ret {
			ret = n
		}
	}
	return
}

func MaxInt(ints ...int) (ret int) {
	ret = math.MinInt
	for _, n := range ints {
		if n > ret {
			ret = n
		}
	}
	return
}
