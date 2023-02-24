package main

// make function to separate digit and append them to new slice
// make new slice in main()
// if not single digit send to function and append returned array to main slice
// if single append it to the main slice

func separator(n int) []int {
	separations := make([]int, 0)
	for n > 0 {
		separations = append([]int{n % 10}, separations...)
		n /= 10
	}
	return separations
}

func separateDigits(nums []int) []int {
	res := make([]int, 0)
	for _, num := range nums {
		if num <= 9 {
			// res = append([]int{num}, res...)
			res = append(res, []int{num}...)
		} else {
			// res = append(separator(num),res...)
			res = append(res, separator(num)...)
		}
	}
	return res
}
