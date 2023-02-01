package backtrack

/*
2
77.组合问题
day:2022-6-16 by lu
update:2023-1-31 by lu
link:https://leetcode.cn/problems/combinations/
question:
	给定两个整数 n 和 k，
	返回 1 ... n 中所有可能的 k 个数的组合。
idea:
	暴力解法（不可靠）：
		有多少k就多少个for循环。
		直接的解法当然是使用for循环，例如示例中k为2，很容易想到 用两个for循环，这样就可以输出 和示例中一样的结果
*/

// 回溯解决组合问题
func combine(n int, k int) [][]int {
	res := [][]int{} // 用来存放符合条件结果的集合
	path := []int{}  // 用来存放符合条件结果
	var backtracking func(start int, path []int)
	backtrack = func(start int, path []int) {
		// base case
		if len(path) == k { // 如果长度达到要求，添加到结果中
			res = append(res, append([]int{}, path...))
			return
		}
		// 回溯算法标准框架
		for i := start; i <= n; i++ { // 剪枝：在每次递归中控制枚举的范围，下一个选择的遍历起点，是当前选择的数字 +1。
			// 选择
			path = append(path, i)
			// 通过start参数控制树枝的遍历，避免产生重复的子集
			backtrack(i+1, path) // 递归往下
			// 撤销选择
			path = path[:len(path)-1] // 回溯
		}
	}
	backtrack(1, path)
	return res
}

// 回溯实现：其实是递归的一种
func combine(n int, k int) [][]int {
	var res [][]int // 符合结果集合
	path := []int{} // 符合结果
	backTrack(1, n, k, path, &res)
	return res
}

func backTrack(start, n, k int, path []int, res *[][]int) { // 注意着里res指针
	// base case
	if len(path) == k {
		*res = append(*res, append([]int{}, path...)) // 注意这里添加结果要重新生成，避免res最终保存时一样的path
		return
	}
	// 回溯算法标准框架
	for i := start; i <= n; i++ {
		path = append(path, i)
		backTrack(i+1, n, k, path, res)
		path = path[:len(path)-1]
	}
}

// 剪枝优化
func combine(n int, k int) [][]int {
	res := [][]int{} // 用来存放符合条件结果的集合
	path := []int{}  // 用来存放符合条件结果
	var backtracking func(start int, path []int)
	backtrack = func(start int, path []int) {
		// base case
		if len(path) == k { // 如果长度达到要求，添加到结果中
			res = append(res, append([]int{}, path...))
			return
		}
		// 回溯算法标准框架
		for i := start; i <= n-(k-len(path)); i++ { // 优化地方
			// 选择
			path = append(path, i)
			// 通过start参数控制树枝的遍历，避免产生重复的子集
			backtrack(i+1, path) // 递归往下
			// 撤销选择
			path = path[:len(path)-1] // 回溯
		}
	}
	backtrack(1, path)
	return res
}
