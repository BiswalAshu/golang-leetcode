package main

// fibbonacci element for that index of n
// ele[1] = 1, ele[2] = 2
func climbStairs(n int) int {
	var (
		result int
		a      int = 1
		b      int = 2
	)
	if n == 1 {
		result = 1
	} else if n == 2 {
		result = 2
	} else {
		for n > 2 {
			result = a + b
			a = b
			b = result
			n--
		}
	}
	return result
}
