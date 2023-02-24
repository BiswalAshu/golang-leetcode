package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func enrich() (primeMap map[byte]int) {
	count := 0
	num := 2
	primeMap = make(map[byte]int)
	for count < 27 {
		if isPrime(num) {
			//add to map here
			primeMap[97+byte(count)] = num
			count++
		}
		num++
	}
	return
}
func stringProduct(str string, primeMap map[byte]int) (prod int64) {
	prod = 1
	for _, i := range str {
		prod *= int64(primeMap[byte(i)])
	}
	return prod
}
func checkInclusion(s1 string, s2 string) bool {

	primeMap := make(map[byte]int)
	primeMap = enrich()

	compareStrProd := stringProduct(s1, primeMap)
	fmt.Println(len(s1))
	fmt.Println()

	for i := 0; i <= (len(s2) - len(s1)); i++ {
		comparisionString := s2[i:(i + len(s1))]
		if stringProduct(comparisionString, primeMap) == compareStrProd {
			return true
		}
	}
	return false
}

//	func copyMap(originalmap map[byte]int) map[byte]int {
//		newMap := make(map[byte]int)
//		for key, val := range originalmap {
//			newMap[key] = val
//		}
//		return newMap
//	}
//
//	func checkInclusion(s1 string, s2 string) bool {
//		chars := make(map[byte]int)
//		for _, letter := range s1 {
//			chars[byte(letter)]++
//		}
//		copyChars := copyMap(chars)
//		for _, letter := range s2 {
//			if copyChars[byte(letter)] > 0 {
//				copyChars[byte(letter)]--
//				if copyChars[byte(letter)] == 0 {
//					delete(copyChars, byte(letter))
//				}
//				if len(copyChars) == 0 {
//					return true
//				}
//			} else {
//				copyChars = copyMap(chars)
//			}
//		}
//		return false
//	}
