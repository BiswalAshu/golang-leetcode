package main

const MOD int64 = 1000000007

func isPrime(n int) bool {
	factorCount := 0
	if n == 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			factorCount++
		}
	}
	if factorCount == 0 {
		return true
	}
	return false
}
func factorial(n int) int {
	fact := 1
	for n > 1 {
		fact = int(int64(fact*n) % MOD)
		n--
	}
	return fact
}
func numPrimeArrangements(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			count++
		}
	}
	res := factorial(n-count) * factorial(count)
	return int(int64(res) % MOD)
}
