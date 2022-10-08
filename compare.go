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

func AllIntsEqual(ints ...int) (yes bool) {
	for i := 1; i < len(ints); i++ {
		if ints[i] != ints[0] {
			return false
		}
	}
	return true
}

func SumInts(ints ...int) (sum int) {
	for _, i := range ints {
		sum += i
	}
	return
}

func IntClamp(x, low, high int) int {
	if x < low {
		x = low
	}
	if x >= high {
		x = high - 1
	}
	return x
}
