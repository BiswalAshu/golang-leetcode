package main

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	sum, current := 10, 9
	for i := 2; i <= n; i++ {
		current *= (9 - i + 2)
		sum += current
	}
	return sum
}
