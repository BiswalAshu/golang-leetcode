package main

func pivotIndex(nums []int) int {
	right := 0
	for _, ele := range nums {
		right += ele
	}
	left := 0
	for ind, ele := range nums {
		// leftSum += ele
		right := right - ele
		if right == left {
			return ind
		}
		left += ele
	}
	return -1
}
