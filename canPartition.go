package main

import "sort"

// func canPartition(nums []int) bool {

// 	if len(nums) < 2 {
// 		return false
// 	}

// 	sort.Ints(nums)
// 	flag:=0
// 	leftSum, rightSum := nums[0], nums[len(nums)-1]
// 	left, right := 0, len(nums)-1
// 	for left < right {
// 		if leftSum < rightSum {
// 			left++
// 			leftSum += nums[left]
// 		} else if rightSum < leftSum {
// 			right--
// 			rightSum += nums[right]
// 		}
// 		if rightSum == leftSum {
// 			flag=1
// 			break
// 		}
// 	}

//		if rightSum == leftSum && flag==0 {
//			return true
//		}
//		return false
//	}
func sumSlice(nums []int) (sum int) {
	sum = 0
	for _, num := range nums {
		sum += num
	}
	return
}

func canPartition(nums []int) bool {
	if len(nums) == 2 && nums[0] == nums[1] {
		return true
	}
	sort.Ints(nums)
	for i := 1; i < len(nums)-2; i++ {
		if sumSlice(nums[:i]) == sumSlice(nums[i+1:]) {
			return true
		}
	}
	return false
}
