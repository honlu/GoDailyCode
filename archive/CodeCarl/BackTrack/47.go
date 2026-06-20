/*
13
47. 全排列2
day: 2022-10-8
update: 2023-2-5 by lu
link: https://leetcode.cn/problems/permutations-ii/
question:
	给定一个可包含重复数字的序列 nums ，
	按任意顺序返回所有不重复的全排列。
answer:
	这道题目和46.全排列 (opens new window)的区别：
	给定的序列有重复数字，要返回所有不重复的全排列。

	注意：去重一定要对元素进行排序，
	这样我们才方便通过相邻的节点来判断是否重复使用了。
*/
// 回溯
func permuteUnique(nums []int) [][]int {
	var res [][]int // 存放结果集
	var path []int  // 存放结果
	sort.Ints(nums) // 注意一定要排序
	// 排列不再使用start做索引参数，而是使用map集合，表示第i位是否使用过
	used := make(map[int]bool)
	// 回溯函数定义
	var backTrack func(used map[int]bool)
	backTrack = func(used map[int]bool) {
		// base case
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		// 回溯标准算法框架，for循环-横向
		for i := 0; i < len(nums); i++ {
			// 去重,防止由于nums重复数字导致的重复排列
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false { // used[i-1]=false是关键
				continue
			}
			if used[i] == true {
				continue // path里已经收录的元素，直接跳过
			}

			path = append(path, nums[i]) // 处理
			used[i] = true               // 处理
			backTrack(used)              // 递归
			used[i] = false              // 回溯
			path = path[:len(path)-1]    // 回溯

		}
	}
	backTrack(used)
	return res
}

// 回溯-拆分
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