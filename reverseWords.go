package main

import "strings"

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func reverseWords(s string) string {
	stringArr := strings.Spilt(s, " ")
	for ind, word := range stringArr {
		stringArr[ind] = reverse(word)
	}
	s = strings.Join(stringArr, " ")
	return s
}
