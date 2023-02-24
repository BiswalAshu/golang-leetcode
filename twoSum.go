package main

func twoSum(numbers []int, target int) []int {
	difference := make(map[int]int)
	result := make([]int, 2)
	for index, num := range numbers {
		diff := target - num
		if val, ok := difference[target-diff]; ok {
			result[0], result[1] = val+1, index+1
			break
		} else {
			difference[diff] = index
		}
	}
	return result
}
