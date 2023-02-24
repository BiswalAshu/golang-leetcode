package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	for index, val := range strs[0] {
		fmt.Print(val)
		for _, word := range strs[1:] {
			// fmt.Println(word)
			if word[index] != byte(val) || index == len(word) {
				return strs[0][:index]
			}
		}
	}
	return strs[0]
}
