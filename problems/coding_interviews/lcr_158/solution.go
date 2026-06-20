package lcr158

import "sort"

// InventoryManagement 返回库存数组中出现次数超过一半的数字。
func InventoryManagement(stock []int) int {
	// 保留原实现思路：排序后取中位数。
	sort.Ints(stock)
	return stock[len(stock)/2]
}
