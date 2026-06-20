/*
209. 长度最小的子数组
2023-1-4
link: https://leetcode.cn/problems/minimum-size-subarray-sum/
question:
	给定一个含有 n 个正整数的数组和一个正整数 target 。

	找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。
	如果不存在符合条件的子数组，返回 0 。

answer:
	双指针实现滑动窗口
*/
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0   // 滑动窗口的起始位置和终止位置
	sum := 0              // 滑动窗口数值之和
	size := math.MaxInt32 // 滑动窗口大小
	for right < len(nums) {
		sum += nums[right]
		for sum >= target { // for循环关键，若使用if就不能一直更新最佳滑动窗口。
			if size > (right - left + 1) {
				size = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++ // 位置要在最后
	}
	if size == math.MaxInt32 { // 没有满足的滑动窗口
		size = 0
	}
	return size
}