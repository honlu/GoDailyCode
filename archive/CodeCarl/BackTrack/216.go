package backtrack

/*
3
216. 组合总和III
day:2022-6-16
update:2023-2-1
link:https://leetcode.cn/problems/combination-sum-iii/
question:
	找出所有相加之和为 n 的 k 个数的组合。
	组合中只允许含有 1 - 9 的正整数，
	并且每种组合中不存在重复的数字。
	返回 所有可能的有效组合的列表 。
	该列表不能包含相同的组合两次，组合可以以任何顺序返回。
idea:
	在[1,2,3,4,5,6,7,8,9]这个集合中找到和为n的k个数的组合。
	本题k相当于了树的深度，9（因为整个集合就是9个数）就是树的宽度。
*/

// 回溯
func combinationSum3(k int, n int) [][]int {
	res := [][]int{} // 符合结果的集合
	path := []int{}  // 符合条件结果
	// 回溯函数定义
	/*
		targetSum: 目标和，即题目中n
		k: 题目中要求k个数的集合
		sum: 已经收集的元素的总和，也就是path里元素的总和
		start: 下一层for循环搜索的起始位置
		(path在参数里加不加都可以)
	*/
	var backTrack func(targetSum, k, sum, start int)
	backTrack = func(targetSum, k, sum, start int) {
		// base case
		if len(path) == k {
			if sum == targetSum {
				res = append(res, append([]int{}, path...)) // 注意这里不能写成res = append(res, path)
			} else {
				return // 如果len(path)==k, 但sum!= targetSum直接返回。
			}
		}
		// 回溯标准算法框架
		for i := start; i <= 9; i++ { // 即递归里面的下一层for循环
			sum += i                          // 处理
			path = append(path, i)            // 处理
			backTrack(targetSum, k, sum, i+1) // 注意i+1调整start
			sum -= i                          // 回溯
			path = path[:len(path)-1]
		}
	}

	backTrack(n, k, 0, 1)
	return res
}

// 回溯+剪枝
/*
剪枝：（剪掉）
	1. 已选元素总和如果大于n,就可以不用往后遍历了
	2. for循环也可以剪枝
*/
func combinationSum3(k int, n int) [][]int {
	res := [][]int{} // 符合结果的集合
	path := []int{}  // 符合条件结果
	// 回溯函数定义
	/*
		targetSum: 目标和，即题目中n
		k: 题目中要求k个数的集合
		sum: 已经收集的元素的总和，也就是path里元素的总和
		start: 下一层for循环搜索的起始位置
	*/
	var backTrack func(targetSum, k, sum, start int)
	backTrack = func(targetSum, k, sum, start int) {
		// base case
		if sum > targetSum { // 剪枝操作
			return // 直接退出
		}
		if len(path) == k {
			if sum == targetSum {
				res = append(res, append([]int{}, path...)) // 注意这里不能写成res = append(res, path)
			} else {
				return // 如果len(path)==k, 但sum!= targetSum直接返回。
			}
		}
		// 回溯标准算法框架
		for i := start; i <= 9-(k-len(path)-1); i++ { // 剪枝
			// 即递归里面的下一层for循环
			sum += i                          // 处理
			path = append(path, i)            // 处理
			backTrack(targetSum, k, sum, i+1) // 注意i+1调整start
			sum -= i                          // 回溯
			path = path[:len(path)-1]
		}
	}

	backTrack(n, k, 0, 1)
	return res
}

// 组合枚举-递归法
func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var temp []int
	var dfs func(cur, rest int)
	dfs = func(cur, rest int) {
		// 找到一个答案
		if len(temp) == k && rest == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		// 剪枝：跳过的数字过多，后面已经无法选到k个数字
		if len(temp)+10-cur < k || rest < 0 {
			return
		}
		// 跳过当前数字
		dfs(cur+1, rest)
		// 选当前数字
		temp = append(temp, cur)
		dfs(cur+1, rest-cur)
		temp = temp[:len(temp)-1]
	}
	dfs(1, n)
	return ans
}
