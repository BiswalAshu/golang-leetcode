package main

func singleNumber(nums []int) int {
	var numCount = make(map[int]int)
	var singleNum int
	for _, i := range nums {
		if _, ok := numCount[i]; ok {
			numCount[i] = 2
		} else {
			numCount[i] = 1
		}
	}
	for k, v := range numCount {
		if v == 1 {
			singleNum = k
		}
	}
	return singleNum
}
