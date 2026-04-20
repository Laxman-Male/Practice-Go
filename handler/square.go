package handler

import "math"

func MyPow(x float64, n int) float64 {
	var result float64
	result = math.Pow(x, float64(n))
	return result
}
