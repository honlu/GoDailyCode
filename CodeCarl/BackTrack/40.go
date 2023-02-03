package backtrack

import "sort"

/*
6
40. 组合总和2
day:2022-10-7
update: 2023-2-3
link: https://leetcode.cn/problems/combination-sum-ii/
question：
	给定一个数组 candidates 和一个目标数 target ，
	找出 candidates 中所有可以使数字和为 target 的组合。
	candidates 中的每个数字在每个组合中只能使用一次。【这里和39不同,以及candidates中元素可能重复】
answer:
	难点：集合（数组candidates）有重复元素，但结果还不能有重复的组合。
	重点：树层去重的话，需要对数组排序！
*/
// 回溯，建议不加path参数，内存消耗更少！这里为了方便理解，回溯函数参数加了path参数。
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int       // 符合结果集
	var path []int        // 符合结果
	sort.Ints(candidates) // 首先排序，方便去重
	// 回溯函数定义和创建:
	/*
		参数定义：
			targetSum:题目要求的目标值
			sum：当前结果路径的和
			start: 下一层for循环搜索的起始位置 [关键理解]]
			path: 可加可不加
	*/
	var backTrack func(targetSum, sum, start int, path []int)
	backTrack = func(targetSum, sum, start int, path []int) {
		// base case
		if sum > targetSum { // 剪枝操作，当然也可以加到下面的for循环条件里。当然是要整个candidate数组有序的情况下。
			return
		}
		if sum == targetSum {
			res = append(res, append([]int{}, path...))
			return
		}
		// 回溯算法标准步骤
		for i := start; i < len(candidates); i++ {
			// 重点在这里
			if i > start && candidates[i] == candidates[i-1] { // 要对同一数层使用过的元素进行跳过
				continue
			}
			sum += candidates[i]                 // 处理
			path = append(path, candidates[i])   // 处理
			backTrack(targetSum, sum, i+1, path) //递归。用i+1,因为每个数字在每个组合中只能使用一次[精华]
			sum -= candidates[i]                 // 回溯
			path = path[:len(path)-1]            // 回溯
		}
	}
	backTrack(target, 0, 0, path)

	return res
}

// 回溯
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	sum := 0
	// 在求和问题中，要去重，先排序是常见的套路！
	sort.Ints(candidates)
	backTrack(0, target, sum, path, candidates, &res)
	return res
}

// 递归函数
func backTrack(startIndex, target, sum int, path, candidates []int, res *[][]int) {
	// 递归的终止条件
	if sum == target {
		temp := make([]int, len(path))
		copy(temp, path) // 深拷贝
		*res = append(*res, temp)
		return
	}
	if sum > target { // 这个条件其实可以省略
		return
	}
	// 剪枝:sum+candidates[i]<=target
	for i := startIndex; i < len(candidates) && sum+candidates[i] <= target; i++ {
		// 同一层去重
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		sum += candidates[i]
		path = append(path, candidates[i]) // 更新路径集合
		// 递归
		backTrack(i+1, target, sum, path, candidates, res)
		// 回溯
		sum -= candidates[i]
		path = path[:len(path)-1]
	}
}
