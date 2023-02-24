package main

import "fmt"

func removeDuplicates(s string) string {
	var stack string
	var head int
	for _, i := range s {
		if len(stack) != 0 && rune(stack[head]) == i {
			head--
		} else {
			stack += string(rune(i))
			head++
		}
	}
	return stack
}
func main() {
	fmt.Print(removeDuplicates("aabbaac"))
}
