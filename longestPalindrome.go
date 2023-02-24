package main

import "fmt"

func longestPalindrome(s string) int {
	charMap := make(map[int]int)
	for _, char := range s {
		charMap[int(char)]++
	}
	oddMax, total := 0, 0
	fmt.Println(charMap)
	for _, v := range charMap {
		if v%2 != 0 && v > oddMax {
			oddMax = v
		}
		if v%2 == 0 {
			total += v
		}
		if v%2 != 0 {
			total += v - 1
		}
	}
	if oddMax != 0 {
		total++
	}
	fmt.Println(charMap)
	return total
}
