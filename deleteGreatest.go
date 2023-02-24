package main

import (
	"sort"
)

func deleteGreatestValue(grid [][]int) int {
	for _, i := range grid {
		sort.Ints(i)
	}
	sum := 0
	colLength := len(grid) - 1
	rowLength := len(grid[0]) - 1
	for j := 0; j < rowLength; j++ {
		max := grid[0][0]
		for i := 0; i < colLength; i++ {
			if grid[i][rowLength] > max {
				max = grid[i][rowLength]
			}
		}
		sum += max

	}
	return sum
}
