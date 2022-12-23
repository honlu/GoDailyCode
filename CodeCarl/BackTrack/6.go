package backtrack

import "sort"

/*
6、组合总和2
day:2022-10-7
link: 20- https://leetcode.cn/problems/combination-sum-ii/
question：给定一个数组 candidates 和一个目标数 target ，
	找出 candidates 中所有可以使数字和为 target 的组合。
	candidates 中的每个数字在每个组合中只能使用一次。【这里和39不同,以及candidates中元素可能重复】
answer:
	难点：集合（数组candidates）有重复元素，但还不能有重复的组合。
	重点：树层去重的话，需要对数组排序！
*/

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
