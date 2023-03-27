package main

func findKthPositive(arr []int, k int) int {
	nums := make(map[int]int)
	for _, v := range arr {
		nums[v] = v
	}
	countIndex := 1
	for i := 1; true; i++ {
		if _, ok := nums[i]; !ok {
			if countIndex == k {
				return i
			} else {
				countIndex++
			}
		}
	}
	return 0
}
