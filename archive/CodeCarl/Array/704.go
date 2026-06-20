/*
704. 二分查找
2023-1-3
link: https://leetcode.cn/problems/binary-search/
question:
	给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，
	写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

answer:
	可以暴力遍历寻找；但有序又不重复数组，最好使用二分查找！
	二分查找关键点是涉及的边界条件，即左右区间的定义。
补充：
	区间的定义是不变量，如果随意些、没有想清楚，很容易二分方法每次写的都不太一样。
	要在二分查找的过程中，保持不变量，就是在while寻找中每一次边界的处理都要坚持根据区间的定义来操作，这就是循环不变量规则。
	区间定义一般有两种：
		1. 左闭右闭即[left,right]
		2. 左闭右开即[left, right)
*/

// 第一种区间定义：左闭右闭
// func search(nums []int, target int) int {
// 	left, right := 0, len(nums)-1
// 	for left <= right { // 左闭右闭区间定义，left=right是有意义的,区间[left,right]依然有效，所以<=
// 		mid := left + (right-left)/2 // 防止溢出 等同于(left + right)/2
// 		if nums[mid] > target {
// 			right = mid - 1 // target 在左区间，所以[left, middle - 1]
// 		} else if nums[mid] < target {
// 			left = mid + 1 // target 在右区间，所以[middle + 1, right]
// 		} else {
// 			return mid // 数组中找到目标值，直接返回下标
// 		}
// 	}
// 	return -1
// }

// 第2种区间定义：左闭右开
func search(nums []int, target int) int {
	left, right := 0, len(nums) // 定义target在左闭右开的区间里，即：[left, right)
	for left < right {          // 因为left = right的时候，在[left, right)是无效的空间，所以使用 <
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid // target 在左区间，在[left, middle)中
		} else if nums[mid] < target {
			left = mid + 1 // target 在右区间，在[middle + 1, right)中
		} else {
			return mid // 数组中找到目标值，直接返回下标
		}
	}
	return -1
}