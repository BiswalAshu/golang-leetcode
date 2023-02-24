package main

import "sort"

func reverseString(s []byte) {
	sort.Reverse(sort.StringSlice(s))
}
