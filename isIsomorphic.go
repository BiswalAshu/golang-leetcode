package main

// rather than mapping e -> g, map letter to their index

func isIsomorphic(s string, t string) bool {
	sMap := make(map[rune]int)
	tMap := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		if sMap[rune(s[i])] != tMap[rune(t[i])] {
			return false
		}
		sMap[rune(s[i])] = i
		tMap[rune(t[i])] = i
	}
	return true
}
