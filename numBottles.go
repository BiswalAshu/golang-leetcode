package main

func numWaterBottles(numBottles int, numExchange int) int {
	var (
		totalDrink int
		remain     int
	)
	totalDrink += numBottles
	for numBottles >= numExchange {
		numBottles = numBottles / numExchange
		remain = numBottles % numExchange
		totalDrink += numBottles
		numBottles = numBottles + remain
	}
	return totalDrink
}
