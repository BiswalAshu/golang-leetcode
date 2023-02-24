package main

func shuffle(nums []int, n int) []int {
	var newArr []int
	for i := 0; i < n; i++ {
		newArr = append(newArr, nums[i])
		newArr = append(newArr, nums[n+i])
	}
	return newArr
}
