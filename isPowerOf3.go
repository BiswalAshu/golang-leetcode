package main

import (
	"math"
)

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	return math.Pow(3, math.Round(math.Log(float64(n))/math.Log(3))) == float64(n)
}
