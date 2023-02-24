package main

func majorityElement(nums []int) int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
		if count[num] > len(nums)/2 {
			return num
		}
	}
	return 0
}
