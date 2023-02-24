package main

func runningSum(nums []int) []int {
	sum := 0
	runArr := make([]int, 0)
	for _, element := range nums {
		sum += element
		runArr = append(runArr, sum)
	}
	return runArr
}
