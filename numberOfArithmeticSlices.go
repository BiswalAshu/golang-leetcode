package main

func isProgressive(nums []int) bool {
	diff := nums[0] - nums[1]
	for i := 1; i < len(nums)-1; i++ {
		if nums[i]-nums[i+1] != diff {
			return false
		}
	}
	return true
}

func numberOfArithmeticSlices(nums []int) int {
	count := 0
	// create subarrays of atleast length 3
	for i := 0; i < len(nums); i++ {
		for j := i + 2; j < len(nums); j++ {
			sub := nums[i : j+1]
			if isProgressive(sub) {
				count++
			}
		}
	}
	return count
}
