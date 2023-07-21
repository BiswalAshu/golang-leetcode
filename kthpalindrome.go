package main

import "math"

func kthPalindrome(queries []int, intLength int) []int64 {
	answerArr := make([]int64, len(queries))

	for i, query := range queries {
		answerArr[i] = getPalindrome(query, intLength)
	}
	return answerArr
}

func getPalindrome(query int, intLengthIn int) int64 {

	if query > int(math.Pow10((intLengthIn+1)/2))-int(math.Pow10((intLengthIn-1)/2)) {
		return int64(-1)
	}

	num := int((query-1)+int(math.Pow10((intLengthIn-1)/2))) * int(math.Pow10((intLengthIn)/2))
	query = (num / (int(math.Pow10((intLengthIn + 1) / 2))))
	count := (intLengthIn) / 2
	revx := 0
	for i := 1; query > 0; i++ {
		revx = revx + (query%10)*int(math.Pow10(count-i))
		query = query / 10
	}

	return int64(num + revx)
}
