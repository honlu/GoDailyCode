package misc

/*
100-287.寻找重复元素
https://leetcode.cn/problems/find-the-duplicate-number/?envType=study-plan-v2&envId=top-100-liked
*/
/*
思路：二分（但是二分的具体实现细节有点难梳理到）

*/
func findDuplicate(nums []int) int {
	// 二分查找
	left, right := 1, len(nums)-1
	res := -1
	for left <= right {
		mid := (left + right) >> 1
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				count++
			}
		}
		if count <= mid {
			left = mid + 1
		} else {
			right = mid - 1
			res = mid
		}
	}
	return res
}
