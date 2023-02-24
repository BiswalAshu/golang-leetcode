package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	s = strings.Trim(s, " ")
	split := strings.Split(s, " ")
	return len(split[len(split)-1])
}
