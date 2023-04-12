/*
442. 数组中重复的数据
2023-4-12 by lu
link: https://leetcode.cn/problems/find-all-duplicates-in-an-array/
question:
	给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，
	且每个整数出现 一次 或 两次 。请你找出所有出现 两次 的整数，并以数组形式返回。

	你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。
answer:
	注意：整数只出现一次或两次。
	方法一：排序+遍历
	方法二：哈希表+遍历
*/
// 方法一：排序
func findDuplicates(nums []int) []int {
	sort.Ints(nums)
	res := []int{}
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] { // 整数只出现一次或两次
			res = append(res, nums[i])
		}
	}
	return res
}

// 方法二：哈希表
func findDuplicates(nums []int) []int {
	hash := make(map[int]bool)
	res := []int{}
	for _, val := range nums {
		if hash[val] == false {
			hash[val] = true
		} else {
			res = append(res, val)
		}
	}
	return res
}