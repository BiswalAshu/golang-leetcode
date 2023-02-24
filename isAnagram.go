package main

func isAnagram(s string, t string) bool {
	letter := make(map[rune]int)
	var (
		flag   int
		result bool
	)
	for _, i := range s {
		if _, ok := letter[i]; ok {
			letter[i]++
		} else {
			letter[i] = 1
		}
	}
	for _, i := range t {
		if _, ok := letter[i]; ok {
			letter[i]--
		} else {
			flag = 1
		}
	}
	for _, val := range letter {
		if val != 0 {
			flag = 1
		}
	}
	if flag == 0 {
		result = true
	} else {
		result = false
	}
	return result
}
