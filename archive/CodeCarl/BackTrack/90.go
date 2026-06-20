/*
10
90. 子集2
day: 2022-10-8
link: https://leetcode.cn/problems/subsets-ii/
question:
	给定一个可能包含重复元素的整数数组 nums，
	返回该数组所有可能的子集（幂集）。
answer:
	和lc-78不一样的地方，数组可能包含重复元素。
	解题框架基本一样，主要加上去重！

	不使用used数组或者set来去重，因为递归的时候下一个start是i+1而不是0。
	如果要是全排列的话，每次要从0开始遍历，为了跳过已入栈的元素，需要使用used。
*/
// 回溯
func subsetsWithDup(nums []int) [][]int {
	var res [][]int // 存结果集
	var path []int  //  为子集收集元素
	sort.Ints(nums) // 避免重复
	// 回溯定义
	var backTrack func(start int)
	backTrack = func(start int) {
		// 每个节点都要添加到结果中
		res = append(res, append([]int{}, path...))

		// 子集好像不需要base case。这个顺序一定要在res添加之后。
		if start == len(nums) {
			return
		}

		// 回溯标准算法步骤
		for i := start; i < len(nums); i++ {
			// 去重
			if i > start && nums[i] == nums[i-1] {
				continue // 不能使用break, 因为{2，2，2，3}后面3还有用
			}
			path = append(path, nums[i]) // 处理
			backTrack(i + 1)             // 递归
			path = path[:len(path)-1]    // 回溯
		}
	}
	backTrack(0)
	return res
}

// 回溯，拆开
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