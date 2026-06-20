package lcr159

import "sort"

// InventoryManagement 返回最少的 cnt 个商品库存数。
func InventoryManagement(stock []int, cnt int) []int {
	// 保留原实现思路：直接排序后截取。
	sort.Ints(stock)
	return stock[:cnt]
}
