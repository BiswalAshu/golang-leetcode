package main

func isSubsequence(s string, t string) bool {
	matchCount := 0
	startFrom := 0
	for _, char := range s {
		for j := startFrom; j < len(t); j++ {
			if byte(char) == t[j] {
				startFrom = j + 1
				matchCount++
				break
			}
		}
	}
	if matchCount == len(s) {
		return true
	}
	return false
}
