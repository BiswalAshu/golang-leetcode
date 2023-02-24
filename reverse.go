package main

func reverse(x int) int {
	sum := 0
	negative := false
	if x < 0 {
		negative = true
		x *= -1
	}
	for x > 0 {
		digit := x % 10
		x /= 10
		sum = (sum * 10) + digit
		if sum > 2147483647 {
			return 0
		}
	}
	if negative == true {
		return sum * -1
	} else {
		return sum
	}
}
