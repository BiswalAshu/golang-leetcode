package main

func moveZeroes(nums []int) {
	next := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[next] = nums[next], nums[i]
			next++
		}
	}
}
