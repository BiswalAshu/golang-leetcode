package main

func selfDividingNumbers(left int, right int) []int {
	var nums []int
	for i := left; i <= right; i++ {
		flag := 0
		backup, n := i, i
		for n > 0 {
			d := n % 10
			n /= 10
			if d == 0 {
				flag = 1
			} else if backup%d != 0 {
				flag = 1
			}
		}
		if flag == 0 {
			nums = append(nums, backup)
		}
	}
	return nums
}
