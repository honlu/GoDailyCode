package main

import "sort"

func inventoryManagement(stock []int) int {
	// 排序
	sort.Ints(stock)
	return stock[(len(stock))/2]
}
