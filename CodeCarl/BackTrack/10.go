/*
10. 子集2
day: 2022-10-8
link: 90- https://leetcode.cn/problems/subsets-ii/
question:
	给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
answer:
	和9基本一样，主要加上去重！
*/
func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	var track []int
	sort.Ints(nums)
	backTrack(0, nums, track, &res)
	return res
}

// 递归
func backTrack(startIndex int, nums, track []int, res *[][]int) {
	// 每个节点都添加到结果中
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp) // 注意这里*

	for i := startIndex; i < len(nums); i++ {
		if i > startIndex && nums[i] == nums[i-1] { // 重复元素排除
			continue
		}
		track = append(track, nums[i])
		backTrack(i+1, nums, track, res)
		// 回溯
		track = track[:len(track)-1]
	}
}