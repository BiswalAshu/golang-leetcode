package main

func checkIfExist(arr []int) bool {
	elementsMap := make(map[int]int)
	for _, num := range arr {
		elementsMap[num]++
		if elementsMap[0] == 2 {
			return true
		}
		if num != 0 {
			if _, ok := elementsMap[num*2]; ok {
				return true
			} else if num%2 == 0 {
				if _, ok := elementsMap[num/2]; ok {
					return true
				}
			}
		}
	}
	return false
}
