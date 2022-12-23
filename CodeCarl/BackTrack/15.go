/*
15. N 皇后
day: 2022-10-9
link: 51-https://leetcode.cn/problems/n-queens/
question:
	n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
	给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
answer:
	难点：二维矩阵
	解决方式：抽象成树来看
	二维矩阵中矩阵的高就是这棵树的高度，
	矩阵的宽就是树形结构中每一个节点的宽度。
	然后：用皇后们的约束条件，来回溯搜索这棵树，只要搜索到了树的叶子节点，
	说明就找到了皇后们的合理位置了。
*/

func solveNQueens(n int) [][]string {
	var res [][]string
	track := make([][]string, n) // 棋盘
	for i := 0; i < n; i++ {
		track[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			track[i][j] = "."
		}
	}
	backTrack(0, n, track, &res)
	return res
}

// 递归 row是当前递归到棋盘的第几行
func backTrack(row, n int, track [][]string, res *[][]string) {
	// 终止条件
	if row == n {
		temp := make([]string, n)
		for k, v := range track { // 将n*n的矩阵，转换成1*n的数组。即将矩阵的一行拼接起来
			temp[k] = strings.Join(v, "")
		}
		*res = append(*res, temp)
		return
	}
	// for 循环 遍历一行，递归遍历列，然后一行一列确定皇后的唯一位置。
	for i := 0; i < n; i++ {
		if isValid(row, i, n, track) { // 验证合法，就看可以防止皇后
			track[row][i] = "Q"
			backTrack(row+1, n, track, res)
			// 回溯,撤销皇后
			track[row][i] = "."
		}
	}
}

// 判断是否符合皇后的条件
func isValid(row, col, n int, track [][]string) bool {
	for i := 0; i < row; i++ { // 检查列。同一列中不能有两个皇后
		if track[i][col] == "Q" {
			return false
		}
	}
	// 检查右斜线，是否有皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if track[i][j] == "Q" {
			return false
		}
	}
	// 检查左斜线，是否有皇后
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if track[i][j] == "Q" {
			return false
		}
	}
	return true
}