/*
13. 全排列2
day: 2022-10-8
link: 47-https://leetcode.cn/problems/permutations-ii/
question:
	给定一个可包含重复数字的序列 nums ，按任意顺序返回所有不重复的全排列。
answer:
	这道题目和46.全排列 (opens new window)的区别
	在与给定一个可包含重复数字的序列，要返回所有不重复的全排列。
注意：去重一定要对元素进行排序，这样我们才方便通过相邻的节点来判断是否重复使用了。
*/
func permuteUnique(nums []int) [][]int {
	var res [][]int
	var track []int
	backTrack(nums, track, &res)
	return res
}

// 递归
func backTrack(nums, track []int, res *[][]int) {
	if len(nums) == 0 {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}
	used := make(map[int]bool) // 唯一的不同，同一层不使用重复的数。
	n := len(nums)
	for i := 0; i < n; i++ {
		if used[nums[i]] == true {
			continue
		}
		cur := nums[i]
		track = append(track, cur)
		used[cur] = true
		// 删除当前元素
		nums = append(nums[:i], nums[i+1:]...)
		backTrack(nums, track, res)
		// 回溯
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		track = track[:len(track)-1]
	}
}