package main

import "sort"

// 返回最少的cnt个商品库存数
func inventoryManagement(stock []int, cnt int) []int {
	// 直接排序算法也可以。
	// 不过题目可能更想考快速排序，后面做一下
	sort.Ints(stock)
	return stock[:cnt]
}
