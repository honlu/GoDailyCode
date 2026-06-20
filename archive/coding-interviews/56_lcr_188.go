package main

import "math"

func bestTiming(prices []int) int {
	cost, profit := math.MaxInt32, 0
	for _, item := range prices {
		cost = min(cost, item)
		profit = max(profit, item-cost)
	}
	return profit
}
