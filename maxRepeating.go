package main

import "strings"

func maxRepeating(sequence string, word string) int {
	k := word
	count := 0
	for strings.Contains(sequence, k) {
		count++
		k += word
	}
	return count
}
