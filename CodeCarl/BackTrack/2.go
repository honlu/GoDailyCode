package backtrack

/*
2.组合问题
day:2022-6-16
link:https://leetcode.cn/problems/combinations/
idea:直接的解法当然是使用for循环，例如示例中k为2，很容易想到 用两个for循环，这样就可以输出 和示例中一样的结果
*/

// 回溯解决组合问题
func combine(n int, k int) [][]int {
	res := [][]int{}
	var backtracking func(start int, path []int)
	backtracking = func(start int, path []int) {
		if len(path) == k { // 如果长度达到要求，添加到结果中
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i <= n; i++ { // 剪枝：在每次递归中控制枚举的范围，下一个选择的遍历起点，是当前选择的数字 +1。
			path = append(path, i)
			backtracking(i+1, path)
			path = path[:len(path)-1] // 回溯
		}
	}
	backtracking(1, []int{})
	return res
}
