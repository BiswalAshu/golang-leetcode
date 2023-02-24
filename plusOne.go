package main

func plusOne(digits []int) []int {
	var new_digits []int

	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	if digits[0] == 0 {
		new_digits = []int{1}
		//return append(new_digits, digits...)
	}
	return append(new_digits, digits...)
}
