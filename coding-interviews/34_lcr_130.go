package main

// LCR 130.衣橱整理
/*
家居整理师将待整理衣橱划分为 m x n 的二维矩阵 grid，其中 grid[i][j] 代表一个需要整理的格子。整理师自 grid[0][0] 开始 逐行逐列 地整理每个格子。

整理规则为：在整理过程中，可以选择 向右移动一格 或 向下移动一格，但不能移动到衣柜之外。同时，不需要整理 digit(i) + digit(j) > cnt 的格子，其中 digit(x) 表示数字 x 的各数位之和。

请返回整理师 总共需要整理多少个格子。（题目不是很好理解，怀疑是最值问题去了）
*/
// 类似：LeetCode 机器人的运动范围
func wardrobeFinishing(m int, n int, cnt int) int {
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 边界，已访问，不需要
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] || (digit(i)+digit(j) > cnt) {
			return 0
		}
		visited[i][j] = true
		// 统计当前格子和向右或向左的格子
		return 1 + dfs(i+1, j) + dfs(i, j+1)
	}
	return dfs(0, 0)
}

func digit(x int) int {
	var res int
	for x > 0 {
		res += x % 10
		x = x / 10
	}
	return res
}
