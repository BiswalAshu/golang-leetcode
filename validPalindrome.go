package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	var res bool
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	newString := reg.ReplaceAllString(s, "")
	newString = strings.ToLower(newString)
	s = ""
	for _, i := range newString {
		s += string(i)
	}
	fmt.Println(newString)
	fmt.Println(s)
	if s == newString {
		res = true
	} else {
		res = false
	}
	return res
}
