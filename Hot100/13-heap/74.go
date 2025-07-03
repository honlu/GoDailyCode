package heap

import "sort"

/*
74-数组中第k个最大元素
*/
func findKthLargest(nums []int, k int) int {
	// 直接调用api，复杂度O(nlogn)，不太符合题目要求复杂度
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// todo: 后续要使用快速排序优化。
