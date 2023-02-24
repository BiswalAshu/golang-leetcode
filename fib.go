package main

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b, count := 0, 1, 2
	sum := 0
	for count <= n {
		count++
		sum = a + b
		a = b
		b = sum
	}
	return sum
}
