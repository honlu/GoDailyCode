/*
题目1020:https://leetcode.cn/problems/number-of-enclaves/description/
2024-1-1
要求：给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。
一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。
返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。

题解：要求找到不靠边的陆地面积。
只要将从周边找到的陆地以及其相邻的陆地，通过BFS或DFS变成海洋。
然后再重新遍历地图的时候，统计剩下的陆地就可以了。
[需要注意遍历细节！]
*/
func numEnclaves(grid [][]int) int {
	// 方向：上、下、左、右
	direct := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	m, n := len(grid), len(grid[0])
	// grid 是地图，也就是一个二维数组
	// x,y 表示开始搜索节点的下标
	var bfs func(grid [][]int, x, y int)
	bfs = func(grid [][]int, x, y int) {
		var queue [][2]int                  // 定义队列
		queue = append(queue, [2]int{x, y}) // 起始节点加入队列
		grid[x][y] = 0                      // 只要加入队列，立刻标记
		for len(queue) != 0 {               // // 开始遍历队列里的元素
			curX, curY := queue[0][0], queue[0][1] // 当前节点坐标
			queue = queue[1:]                      // 从队列取元素
			for i := 0; i < 4; i++ {               // 开始当前节点的四个方向上下左右遍历
				nextX := curX + direct[i][0]
				nextY := curY + direct[i][1]
				if nextX < 0 || nextX >= m || nextY < 0 || nextY >= n {
					continue // 坐标越界了，直接跳过
				}
				if grid[nextX][nextY] == 1 { // 如果节点是陆地
					queue = append(queue, [2]int{nextX, nextY}) // 队列添加该节点为下一轮要遍历的节点
					grid[nextX][nextY] = 0                      // 只要加入队列立刻标记，避免重复访问
				}
			}
		}
	}

	// 从左侧边、右侧边向中间遍历
	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			bfs(grid, i, 0)
		}
		if grid[i][n-1] == 1 {
			bfs(grid, i, n-1)
		}
	}
	// 从上、下边，向中间遍历
	for i := 0; i < n; i++ {
		if grid[0][i] == 1 {
			bfs(grid, 0, i)
		}
		if grid[m-1][i] == 1 {
			bfs(grid, m-1, i)
		}
	}
	// 以上结束。剩下的就是不靠边的陆地面积
	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				res++
			}
		}
	}
	return res
}