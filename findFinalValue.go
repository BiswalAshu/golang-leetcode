package main

func findFinalValue(nums []int, original int) int {
	elements := make(map[int]int)

	if original == 0 {
		return original
	}
	for _, num := range nums {
		elements[num] = num
	}

	for {
		if val, ok := elements[original]; ok {
			original = val * 2
		} else {
			break
		}
	}
	return original
}
