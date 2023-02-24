package main

func rotate(nums []int, k int) {
	idx := len(nums) - (k % len(nums))
	newNums := append(nums[idx:], nums[:idx]...)
	copy(nums, newNums)
}
