package main

func dominantIndex(nums []int) int {
	max, maxInd := nums[0], 0
	for ind, num := range nums {
		if num > max {
			max = num
			maxInd = ind
		}
	}
	for _, num := range nums {
		if num*2 > max && num != max {
			return -1
		}
	}
	return maxInd
}

// func dominantIndex(nums []int) int {
// 	max, maxInd := nums[0], 0
// 	elements := make(map[int]int)
// 	for ind, num := range nums {
// 		if num > max {
// 			max = num
// 			maxInd = ind
// 		}
// 	}
// 	for index, num := range nums {
// 		if num != max {
// 			elements[num*2] = index
// 		}
// 	}
// 	for key,_ := range elements {
// 		if key > max {
// 			return -1
// 		}
// 	}
// 	return maxInd
// }
