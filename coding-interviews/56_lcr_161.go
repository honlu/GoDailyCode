package main

import "math"

func maxSales(sales []int) int {
	if len(sales) == 0 {
		return 0
	}
	var tempMax int
	var res int = math.MinInt32
	for i := 0; i < len(sales); i++ {
		tempMax = max(tempMax+sales[i], sales[i])
		res = max(tempMax, res)
	}
	return res
}
