/*
12. 全排列
day： 2022-10-8
link: 46-https://leetcode.cn/problems/permutations/
question:
	给定一个 没有重复 数字的序列，返回其所有可能的全排列。
answer:
	首先排列是有序的，也就是说 [1,2] 和 [2,1] 是两个集合，
	这和之前分析的子集以及组合所不同的地方。

注意：单层搜索的逻辑
这里和77.组合问题 (opens new window)、131.切割问题 (opens new window)和78.子集问题 (opens new window)
最大的不同就是for循环里不用startIndex了。
因为排列问题，每次都要从头开始搜索，例如元素1在[1,2]中已经使用过了，
但是在[2,1]中还要再使用一次1。
而used数组，其实就是记录此时path里都有哪些元素使用了，
一个排列里一个元素只能使用一次。
*/
func permute(nums []int) [][]int {
	var res [][]int
	var track []int
	backTrack(nums, track, &res)
	return res
}

// 递归-纵向循环
func backTrack(nums, track []int, res *[][]int) {
	// 符合条件
	if len(nums) == 0 {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
	}
	// for 循环-横向
	n := len(nums)
	for i := 0; i < n; i++ {
		cur := nums[i]
		track = append(track, cur)
		nums = append(nums[:i], nums[i+1:]...) // 删除当前元素
		backTrack(nums, track, res)
		// 回溯
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...) // 回复原来的数组
		track = track[:len(track)-1]
	}
}