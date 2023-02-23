/*
11
491. 递归子序列
day: 2022-10-8
update: 2023-2-5 by lu
link: https://leetcode.cn/problems/increasing-subsequences/
question:
	给定一个整型数组, 你的任务是找到所有该数组的递增子序列，
	递增子序列的长度至少是2。
answer:
	注意：递增子序列长度至少要是2,所以通过树来看，
	节点的序列长度至少是2切实递增才能添加到结果里。
	另外：同一父节点下的同层上使用过的元素就不能再使用了。使用map去重。
*/
// 回溯
func findSubsequences(nums []int) [][]int {
	var res [][]int // 符合结果的集合
	var path []int  // 符合结果的路径
	// 回溯函数定义和申明
	var backTrack func(start int)
	backTrack = func(start int) {
		// base case
		if len(path) >= 2 { // 至少两个元素
			res = append(res, append([]int{}, path...))
		}
		// 回溯算法标准逻辑框架
		used := make(map[int]bool) // 同一父节点下的同层元素是否使用过。
		for i := start; i < len(nums); i++ {
			if used[nums[i]] == true { // (关键)前面已经用过就跳过
				continue
			}
			if len(path) > 0 && nums[i] < path[len(path)-1] { // 当前元素不符合递增条件，跳过
				continue
			}
			// 处理
			used[nums[i]] = true
			path = append(path, nums[i])
			backTrack(i + 1) // 递归
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}

// 回溯-拆开写
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