package main

func indexOf(data []int, element int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	indexList := make([]int, 0)
	for _, num1 := range nums1 {
		foundAt := indexOf(nums2, num1)
		// fmt.Println(foundAt)
		max := num1
		flag := 0
		for _, num2 := range nums2[foundAt:] {
			if num2 > max {
				// append to slice
				indexList = append(indexList, num2)
				flag = 1
				break
			}

		}
		if flag == 0 {
			indexList = append(indexList, -1)
		}
	}
	return indexList
}
