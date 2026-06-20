/*
15
51. N 皇后
day: 2022-10-9
update: 2023-2-6
link: https://leetcode.cn/problems/n-queens/
question:
	n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，
	并且使皇后彼此之间不能相互攻击。
	 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
answer:
	难点：二维矩阵
	解决方式：抽象成树来看
	二维矩阵中矩阵的高就是这棵树的高度，矩阵的宽就是树形结构中每一个节点的宽度。
	然后：用皇后们的约束条件，来回溯搜索这棵树，
	只要搜索到了树的叶子节点，说明就找到了皇后们的合理位置了。
*/
// 第一版回溯
func solveNQueens(n int) [][]string {
	var res [][]string  // 符合结果的集合
	var path [][]string // 符合结果,预计下面初始化
	path = make([][]string, n)
	for i := 0; i < n; i++ {
		path[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			path[i][j] = "."
		}
	}
	// 回溯函数定义,参数row,定义递归层数，也是棋盘的第几行
	var backTrack func(row int)
	backTrack = func(row int) {
		// base case
		if row == n {
			var temp []string
			for i := 0; i < n; i++ {
				temp = append(temp, strings.Join(path[i], ""))
			}
			res = append(res, temp)
			return
		}
		// 回溯标准框架，for循环-横向
		for i := 0; i < n; i++ {
			if isValid(row, i, n, path) { // 在[row,i]插入皇后，看是否合法
				path[row][i] = "Q" // 处理
				backTrack(row + 1) // 递归
				path[row][i] = "." // 回溯
			}
		}
	}

	backTrack(0)
	return res
}

// 判断是否符合n皇后的条件
func isValid(row, col, n int, path [][]string) bool {
	// 检查列，在之前是否已经有皇后.这是一个剪枝
	for i := 0; i < row; i++ {
		if path[i][col] == "Q" {
			return false
		}
	}
	// 检查135度，右斜线，在之前是否已经有皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if path[i][j] == "Q" {
			return false
		}
	}
	// 检查45度， 左斜线， 在之前是否已经有皇后
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 { // 注意j初始值col+1
		if path[i][j] == "Q" {
			return false
		}
	}
	return true
}

// 第二版：回溯分开写
// func solveNQueens(n int) [][]string {
// 	var res [][]string
// 	track := make([][]string, n) // 棋盘
// 	for i := 0; i < n; i++ {
// 		track[i] = make([]string, n)
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			track[i][j] = "."
// 		}
// 	}
// 	backTrack(0, n, track, &res)
// 	return res
// }

// // 递归 row是当前递归到棋盘的第几行
// func backTrack(row, n int, track [][]string, res *[][]string) {
// 	// 终止条件
// 	if row == n {
// 		temp := make([]string, n)
// 		for k, v := range track { // 将n*n的矩阵，转换成1*n的数组。
// 			// 即将矩阵的一行拼接起来
// 			temp[k] = strings.Join(v, "")
// 		}
// 		*res = append(*res, temp)
// 		return
// 	}
// 	// for 循环 遍历一行，递归遍历列，然后一行一列确定皇后的唯一位置。
// 	for i := 0; i < n; i++ {
// 		if isValid(row, i, n, track) { // 验证合法，就看可以防止皇后
// 			track[row][i] = "Q"
// 			backTrack(row+1, n, track, res)
// 			// 回溯,撤销皇后
// 			track[row][i] = "."
// 		}
// 	}
// }

// // 判断是否符合皇后的条件（不能同行，不能同列，不能同斜度（45度和135度角））
// /*
// 为什么没有在同行进行检查呢？
// 因为在单层搜索的过程中，每一层递归，
// 只会选for循环（也就是同一行）里的一个元素，所以不用去重了。
// */
// func isValid(row, col, n int, track [][]string) bool {
// 	for i := 0; i < row; i++ { // 检查列。同一列中不能有两个皇后
// 		if track[i][col] == "Q" {
// 			return false
// 		}
// 	}
// 	// 检查右斜线，是否有皇后
// 	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
// 		if track[i][j] == "Q" {
// 			return false
// 		}
// 	}
// 	// 检查左斜线，是否有皇后
// 	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
// 		if track[i][j] == "Q" {
// 			return false
// 		}
// 	}
// 	return true
// }