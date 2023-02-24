package main

import "sort"

func sortedSquares(nums []int) []int {
	// sortedArr := make([]int, len(nums))
	for i, num := range nums {
		nums[i] = num * num
	}
	sort.Ints(nums)
	return nums
}
