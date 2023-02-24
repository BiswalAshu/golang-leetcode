package main

import "fmt"

func isValid(s string) bool {
	var stack []rune
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
		case '{':
		case '[':
			stack = append(stack, rune(i))

			break
		case ')':
			if len(stack) != 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]

				return false
			}
			break
		case '}':
			if len(stack) != 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
				fmt.Println(stack)
				return false
			}
			break
		case ']':
			if len(stack) != 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
				return false
			}
			break
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}
