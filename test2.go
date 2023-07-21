// 2 3 4 5 6 9 8 6

// k = 10

// 2 3 5 7 5

// count of subset of list1 should b k as sum

package main

import "fmt"

func subsetSum(arr []int, k int) int {
	subsetCount := 0
	for _, i := range arr {
		sum := 0
		for j := i; i < len(arr); i++ {
			if sum == k {
				subsetCount++
				// break
			} else if k >= sum+arr[j] {
				sum += arr[j]
			} else if sum+arr[j] < k {
				sum += arr[j]
				// continue
			}
		}
	}
	return subsetCount
}
func main() {
	// list1 := make([]int,15)
	list1 := []int{2, 3, 5, 7, 5}
	k := 10

	result := subsetSum(list1, k)
	fmt.Println(result)
}
