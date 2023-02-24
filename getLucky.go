package main

import (
	"fmt"
)

func charToInt(s string) (res int) {
	res = 0
	for _, char := range s {
		if char-96 <= 9 {
			res = (res * 10) + int(char-96)
		} else {
			res = (res * 100) + int(char-96)
		}
	}
	fmt.Println(res)
	return
}
func getLucky(s string, k int) int {
	num := charToInt(s)
	for k > 0 {
		sum := 0
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		num = sum
		k--
	}
	return num
}
