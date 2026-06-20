/*
977. 有序数组的平方
2023-1-3
link： https://leetcode.cn/problems/squares-of-a-sorted-array/
question:
	给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，
	要求也按 非递减顺序 排序。
answer:
	暴力解法：先每个树平方得到结果数组，然后排序

	双指针法：定义一个结果数组，然后最后保存平方和最大值。
*/
package main

// 暴力解法 时间和空间复杂度
// func sortedSquares(nums []int) []int {
// 	var temp []int
// 	for _, val := range nums {
// 		temp = append(temp, val*val)
// 	}
// 	sort.Ints(temp)
// 	return temp
// }
// func sortedSquares(nums []int) []int {
// 	for i, val := range nums {
// 		nums[i] = val * val
// 	}
// 	sort.Ints(nums)
// 	return nums
// }

// 双指针：规律性总结
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	size := len(nums) - 1
	left, right := 0, len(nums)-1 // 注意区间定义，左闭右闭
	for left <= right {           // 区间定义，left = right有意义
		if nums[left]*nums[left] < nums[right]*nums[right] {
			res[size] = nums[right] * nums[right]
			right--
			size--
		} else {
			res[size] = nums[left] * nums[left]
			left++
			size--
		}
	}
	return res
}
