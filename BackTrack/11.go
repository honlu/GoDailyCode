/*
11. 递归子序列
day: 2022-10-8
link: 491-https://leetcode.cn/problems/increasing-subsequences/
question:
	给定一个整型数组, 你的任务是找到所有该数组的递增子序列，递增子序列的长度至少是2。
answer:
*/
func findSubsequences(nums []int) [][]int {
	var track []int
	var res [][]int
	backTrack(0, nums, track, &res)
	return res
}

// 递归
func backTrack(startIndex int, nums, track []int, res *[][]int) {
	if len(track) > 1 { // 符合条件，至少有两个元素
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}
	used := make(map[int]bool) // 每层都新建一个哈希表用于去重
	for i := startIndex; i < len(nums); i++ {
		// 分两种情况判断：一：当前取的元素小于子集的最后一个元素，则继续寻找下一个适合的元素
		// 二：当前取的元素在本层已经出现过，所以跳过该元素，继续寻找
		if len(track) > 0 && nums[i] < track[len(track)-1] || used[nums[i]] == true {
			continue
		}
		used[nums[i]] = true // 表示本层该元素已经使用过了
		track = append(track, nums[i])
		backTrack(i+1, nums, track, res)
		// 回溯
		track = track[:len(track)-1]
	}
}