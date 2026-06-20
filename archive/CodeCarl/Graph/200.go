/*
题目：https://leetcode.cn/problems/number-of-islands/description/

要求：给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。

题解：

*/

func numIslands(grid [][]byte) int {
	// 方向：上、下、左、右
	direct := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	// grid 是地图，也就是一个二维数组
	// visited标记访问过的节点，不要重复访问
	// x,y 表示开始搜索节点的下标
	var bfs func(grid [][]byte, visited [][]bool, x, y int)
	bfs = func(grid [][]byte, visited [][]bool, x, y int) {
		var queue [][2]int                  // 定义队列
		queue = append(queue, [2]int{x, y}) // 起始节点加入队列
		visited[x][y] = true                // 只要加入队列，立刻标记为访问过的节点
		for len(queue) != 0 {               // // 开始遍历队列里的元素
			curX, curY := queue[0][0], queue[0][1] // 当前节点坐标
			queue = queue[1:]                      // 从队列取元素
			for i := 0; i < 4; i++ {               // 开始当前节点的四个方向上下左右遍历
				nextX := curX + direct[i][0]
				nextY := curY + direct[i][1]
				if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n {
					continue // 坐标越界了，直接跳过
				}
				if !visited[nextX][nextY] && grid[nextX][nextY] == '1' { // 如果节点没被访问过
					queue = append(queue, [2]int{nextX, nextY}) // 队列添加该节点为下一轮要遍历的节点
					visited[nextX][nextY] = true                // 只要加入队列立刻标记，避免重复访问
				}

			}
		}
	}
	// 开始看岛屿数量
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				res++                    // 遇到没访问过的陆地，+1
				bfs(grid, visited, i, j) // 将与其链接的陆地都标记上 true
			}
		}
	}
	return res
}