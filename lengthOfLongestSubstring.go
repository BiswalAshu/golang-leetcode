package main

// func lengthOfLongestSubstring(s string) int {
// 	max, end, length := 0, 0, 0
// 	charExists := make(map[byte]int)
// 	for start := 0; start < len(s); start++ {
// 		if _, ok := charExists[s[start]]; !ok {
// 			charExists[s[start]]++
// 			length = start - end + 1
// 			fmt.Print("char added --")
// 			fmt.Print(rune(s[start]))
// 			fmt.Println(charExists)
// 		} else {
// 			end++
// 			for s[end] != s[start] {
// 				delete(charExists, s[end])
// 				fmt.Print("char delete --")
// 				fmt.Print(rune(s[end]))
// 				fmt.Println(charExists)
// 				end++
// 			}
// 		}
// 		if max < length {
// 			max = length
// 			fmt.Print("MAXXX. ")
// 			fmt.Println(max)
// 		}
// 	}
// 	return max
// }

// func lengthOfLongesSubstring(s string) int {
// 	result := 0
// 	seen := map[byte]bool{}
// 	start := 0
// 	for end := 0; end < len(s); end++ {

// 		for seen[s[end]] {
// 			// remove whats at start, and "pinch" the sliding window
// 			seen[s[start]] = false
// 			start++
// 		}

// 		// we can put this one here now, its safe
// 		seen[s[end]] = true

// 		// now here, we are guaranteed to have a "window" with no repeaters
// 		if end-start+1 > result {
// 			result = end - start + 1
// 		}
// 	}

// 	return result
// }

func lengthOfLongestSubstring(s string) int {
	charExists := make(map[rune]int)
	maxLength, end := 0, 0
	for start, char := range s {
		if index, ok := charExists[char]; ok && index >= end {
			end = index + 1
		}
		charExists[char] = start
		currentLength := start - end + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}
