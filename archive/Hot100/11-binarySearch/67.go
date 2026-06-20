package binarysearch

/*
67-153.寻找旋转排序数组中的最小值
https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/?envType=study-plan-v2&envId=top-100-liked
*/

/*
方法一：从中间点开始。时间复杂度从O(n)转为O(n/2)
*/
// func findMin(nums []int) int {
// 	left, right := 0, len(nums)-1
// 	if left < right {
// 		mid := left + right>>1
// 		if nums[left] < nums[mid] {
// 			left = mid - 1 // 很好处理case类型：[3,1,2]
// 		}
// 		if nums[mid] < nums[right] {
// 			right = mid // 很好
// 		}
// 		for i := left + 1; i <= right; i++ {
// 			if nums[i] < nums[i-1] {
// 				return nums[i]
// 			}
// 		}
// 	}
// 	return nums[0] // 很好处理[1],[11,13,15,17]case
// }
/*
方法2：时间复杂度转为O(log(n))
*/
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1 // 这里关键
		}
	}
	return nums[left]
}
