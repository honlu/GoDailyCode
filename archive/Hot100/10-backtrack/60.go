package backtrack

/*
60-79.单词搜索
https://leetcode.cn/problems/word-search/?envType=study-plan-v2&envId=top-100-liked
*/
/*
思路：DFS+回溯
*/
func exist(board [][]byte, word string) bool {
	// 深度优先：即暴力遍历矩阵中所有以（i,j）位置出发，是否能搜到word单词。
	// DFS是通过递归，先朝一个方向搜到底，再回溯至上个节点，沿另外一个方向搜索查看。
	// 从每个元素出发，有四个方向可以选择：上下左右
	// 在DFS同时，遍历word字符串对比验证，如果相等，则继续DFS验证；如果不等，则停止DFS回溯。当遍历到字符串word末尾，则存在
	x := [4]int{-1, 0, 1, 0}
	y := [4]int{0, 1, 0, -1}
	m, n := len(board), len(board[0])
	var res bool
	var dfs func(i, j, k int)
	dfs = func(i, j, k int) {
		if i < 0 || j < 0 || i >= m || j >= n || res {
			return
		}
		// 剪枝
		if board[i][j] == '*' || board[i][j] != word[k] { // 走过的不再走|| 不匹配，就回去，重新DFS
			return
		}
		if k == len(word)-1 {
			res = true
			return
		}
		temp := board[i][j]
		board[i][j] = '*' // 标记此节点已经走过
		// dfs下一个元素
		for o := 0; o < 4; o++ {
			dfs(i+x[o], j+y[o], k+1)
		}
		// 回溯
		board[i][j] = temp
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, 0)
			if res {
				return res
			}
		}
	}
	return res
}
