package main

func subsets(nums []int) [][]int {
	var res [][]int
	var curr []int
	backtrack(nums, &res, curr, 0)
	return res
}

func backtrack(nums []int, res *[][]int, curr []int, start int) {
	tmp := make([]int, len(curr))
	copy(tmp, curr)
	*res = append(*res, tmp)
	for i := start; i < len(nums); i++ {
		curr = append(curr, nums[i])
		backtrack(nums, res, curr, i+1)
		curr = curr[:len(curr)-1]
	}
}
