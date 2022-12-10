/*
695. 岛屿的最大面积
2022-12-10
link: https://leetcode.cn/problems/max-area-of-island/
question:
	给你一个大小为 m x n 的二进制矩阵 grid 。
	岛屿 是由一些相邻的 1 (代表土地) 构成的组合，
	这里的「相邻」要求两个 1 必须在 水平或者竖直的四个方向上 相邻。
	你可以假设 grid 的四个边缘都被 0（代表水）包围着。
	岛屿的面积是岛上值为 1 的单元格的数目。
	计算并返回 grid 中最大的岛屿面积。如果没有岛屿，则返回面积为 0 。
answer:
	相邻、水平或竖直，则发现是搜索性问题！
	这里采用深度优先算法
*/
type pair struct {
	x, y int
}

var dirs = []pair{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // 行，列

func maxAreaOfIsland(grid [][]int) int {
	var res = 0
	for i, row := range grid { // 双层循环遍历+深度优先搜索.每个节点都要深度优先搜索
		for j, _ := range row {
			temp := dfs(grid, i, j)
			if res < temp {
				res = temp
			}
		}
	}
	return res
}

func dfs(grid [][]int, x, y int) int {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] != 1 {
		return 0
	}
	grid[x][y] = 0 // 土地访问不会超过一次，设为0则可以实现
	var res = 1
	// 深度优先搜索
	for _, dir := range dirs {
		res += dfs(grid, x+dir.x, y+dir.y) // 找准一个方向深度递归
	}

	return res
}