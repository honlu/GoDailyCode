/*
163. 寻找峰值
2022-12-9
link: https://leetcode.cn/problems/find-peak-element/
question:
	峰值元素是指其值严格大于左右相邻值的元素。
	给你一个整数数组 nums，找到峰值元素并返回其索引。
	数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。

answer:
	要求实现时间复杂度为 O(log n) 的算法来解决此问题。
	所以简单遍历会超出时间限制！要巧妙一些,使用二分查找。
*/
func findPeakElement(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > nums[mid+1] { // 很巧秒的判断和移动
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}