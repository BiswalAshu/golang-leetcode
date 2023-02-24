package main

func searchInsert(nums []int, target int) int {
	pos := 0
	if nums[0] > target {
		pos = 0
	} else if target > nums[len(nums)-1] {
		pos = len(nums)
	} else {
		for ind, val := range nums {
			if val == target {
				pos = ind
			}
			if val < target && target < nums[ind+1] {
				pos = ind + 1
			}
		}
	}
	return pos
}
