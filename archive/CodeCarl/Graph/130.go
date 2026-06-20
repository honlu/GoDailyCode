/*
题目：130.被围绕的区域
	https://leetcode.cn/problems/surrounded-regions/description/
要求：
	给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，
	找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
题解：
	被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。
	如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
	【反向思维】
*/
func solve(board [][]byte) {
	// 方向：上下左右
	direct := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(board), len(board[0])
	// board是矩阵，即二维矩阵
	// x,y 表示开始搜索节点的下标。 广度有限搜索
	var bfs func(board [][]byte, x, y int)
	bfs = func(board [][]byte, x, y int) {
		var queue [][2]int                  // 定义队列
		queue = append(queue, [2]int{x, y}) // 起始节点加入队列
		board[x][y] = 'A'                   // 只要加入队列，‘0’立刻改为‘A’
		for len(queue) != 0 {               // 开始遍历队列中的元素，按照广度
			curX, curY := queue[0][0], queue[0][1]
			queue = queue[1:] // 出队列
			for i := 0; i < 4; i++ {
				nextX := curX + direct[i][0]
				nextY := curY + direct[i][1]
				if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n {
					continue // 坐标越界，跳过
				}
				if board[nextX][nextY] == 'O' { // 与边界O相邻的O
					queue = append(queue, [2]int{nextX, nextY})
					board[nextX][nextY] = 'A' // 只要是O加入队列就要改为A
				}
			}
		}
	}
	// 从左右两侧符合O的遍历
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			bfs(board, i, 0)
		}
		if board[i][n-1] == 'O' {
			bfs(board, i, n-1)
		}
	}
	// 从上下边符合O的遍历
	for i := 0; i < n; i++ {
		if board[0][i] == 'O' {
			bfs(board, 0, i)
		}
		if board[m-1][i] == 'O' {
			bfs(board, m-1, i)
		}
	}
	// 当与周边所处的O变成A时，内部被包围的O就可以变成X。周边的A改会O
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}