package binarysearch

/*
66-33. 搜索旋转排序数组
https://leetcode.cn/problems/search-in-rotated-sorted-array/description/?envType=study-plan-v2&envId=top-100-liked
*/
func search(nums []int, target int) int {
	// 哈希处理
	hash := map[int]int{}
	for i, val := range nums {
		hash[val] = i
	}
	if hash[target] != 0 || nums[0] == target {
		return hash[target]
	}
	return -1
}
