package backtrack

import "sort"

/*
78-子集
https://leetcode.cn/problems/subsets/?envType=study-plan-v2&envId=top-100-liked
*/
func subsets(nums []int) [][]int {
	// 子集问题是从哪个位置选的问题，不是用不用的问题
	sort.Ints(nums) // 先排个序
	var res [][]int
	var path []int
	var recursion func(index int)
	recursion = func(index int) {
		res = append(res, append([]int{}, path...))
		if len(path) == len(nums) {
			return
		}
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			recursion(i + 1)
			path = path[:len(path)-1] // 回溯
		}
	}
	recursion(0)
	return res
}
