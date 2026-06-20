package binarysearch

/*
63-搜索插入位置
https://leetcode.cn/problems/search-insert-position/?envType=study-plan-v2&envId=top-100-liked
*/
// 二分算法求解
func searchInsert(nums []int, target int) int {
	// 怎么避免超时
	low, higth := 0, len(nums)
	for low < higth {
		mid := (low + higth) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			higth = mid
		} else {
			low = mid + 1 // 注意这里+1细节，否则会超时
		}
	}
	return low
}

// 暴力解法
func searchInsertV(nums []int, target int) int {
	// 怎么避免超时
	for i, item := range nums {
		if item < target {
			continue
		} else {
			return i
		}
	}
	return len(nums)
}
