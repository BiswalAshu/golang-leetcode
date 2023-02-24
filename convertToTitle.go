package main

func convertToTitle(columnNumber int) string {
	result := ""
	for columnNumber > 0 {
		remainder := columnNumber % 26
		if remainder == 0 {
			remainder = 26
		}
		result = string(rune(remainder+64)) + result
		columnNumber = (columnNumber - remainder) / 26
	}
	return result
}
