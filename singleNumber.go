package main

func singleNumber(nums []int) int {
	countMap := make(map[int]int)
	for _, val := range nums {
		countMap[val]++
		if countMap[val] == 3 {
			delete(countMap, val)
		}
	}
	// fmt.Println(countMap)
	key := 0
	for k := range countMap {
		key = k
		break
	}
	return key
}
