package main

func countConsistentStrings(allowed string, words []string) int {
	allowMap := make(map[rune]rune)
	for _, i := range allowed {
		if _, ok := allowMap[i]; !ok {
			allowMap[i] = i
		}
	}
	// fmt.Println(allowMap)
	notConsitentCount := 0
	for _, word := range words {
		// fmt.Println(word)
		for _, ch := range word {
			if _, ok := allowMap[ch]; !ok {
				notConsitentCount++
				break
			}
		}
	}
	return len(words) - notConsitentCount
}
