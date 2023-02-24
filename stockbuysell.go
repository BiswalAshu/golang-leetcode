package main

func maxProfit(prices []int) int {
	min, max := prices[0], prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < min {
			min = price
			max = price
		}
		if price > max {
			max = price
		}
		profit := max - min
		if profit > maxProfit {
			maxProfit = profit
		}

	}
	return maxProfit
}
