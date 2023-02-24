package main

func isHappy(n int) bool {
	sum := 0
	var track []int
	for n >= 9 {
		for n > 0 {
			d := n % 10
			sum = sum + (d * d)
			n = n / 10
		}
		n = sum
		track = append(track, sum)

	}
	if n == 1 {
		return true
	}
	return false
}
