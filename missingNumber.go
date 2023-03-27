package main

func missingNumber(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	total := (len(nums) * (len(nums) + 1)) / 2
	return total - sum
}
