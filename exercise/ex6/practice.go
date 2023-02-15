package ex6

import "math"

func MinInt(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}
