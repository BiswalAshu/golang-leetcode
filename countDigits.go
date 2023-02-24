package main

func countDigits(num int) int {
	var count int
	backup := num
	for num > 0 {
		d := num % 10
		num /= 10
		if backup%d == 0 {
			count++
		}
	}
	return count
}
