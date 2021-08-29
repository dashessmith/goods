package goods

import "math"

func RoundFloat64(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func FixedFloat64(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(RoundFloat64(num*output)) / output
}
