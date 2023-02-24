package main

import "fmt"

func maxNumberOfBalloons(text string) int {
	balChars := make(map[rune]int)
	for _, char := range text {
		// fmt.Println(char)
		if rune(char) == 'b' || rune(char) == 'a' || rune(char) == 'l' || rune(char) == 'o' || rune(char) == 'n' {
			balChars[rune(char)]++
		}
	}
	count := 0
	for balChars['b'] >= 1 && balChars['a'] >= 1 && balChars['l'] >= 2 && balChars['o'] >= 2 && balChars['n'] > 1 {
		count++
		balChars['b']--
		balChars['a']--
		balChars['l'] -= 2
		balChars['o'] -= 2
		balChars['n']--
	}
	fmt.Println(balChars)
	return count
}
