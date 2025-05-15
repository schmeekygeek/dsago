package arrays

import "math"

func MaxProfit(prices []int) int {
	l := 0
	r := 1
	maxProfit := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			maxProfit = int(math.Max(float64(maxProfit), float64(profit)))
		} else {
			l = r
		}
		r++
	}

	return maxProfit
}
