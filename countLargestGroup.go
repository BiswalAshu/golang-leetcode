package main

import "fmt"

func addDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10)
		n /= 10
	}
	return sum
}

// group[ value of added digits ] = number of such groups

func countLargestGroup(n int) int {
	groups := make(map[int]int)
	for i := 1; i <= n; i++ {
		sum := addDigits(i)
		groups[sum]++
	}
	// findin max value
	max := 0
	for _, val := range groups {
		if val > max {
			max = val
		}
	}
	count := 0
	//count repetation of max
	for _, val := range groups {
		if val == max {
			count++
		}
	}
	fmt.Println(groups)
	return count
}
