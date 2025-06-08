package hot100

/*
52-腐烂的橘子
https://leetcode.cn/problems/rotting-oranges/description/?envType=study-plan-v2&envId=top-100-liked

橘子腐烂过程和广搜层次差不多。
算法解题模型转化：多源最短路径。（dfs没有多源，bfs多源就是指多个起点）
新鲜橘子到腐烂橘子的最近距离。
*/
func orangesRotting(grid [][]int) int {
	// 类似广度优先遍历深度，但是有多个「腐烂的橘子」起点
	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}
	var queue [][2]int
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	var res int // 注意特殊情况：0分钟时就没有新鲜橘子
	if len(queue) > 0 {
		res-- // 分钟数比层数少一
	}
	// 多层遍历，并且多个起点
	for len(queue) > 0 {
		res++
		size := len(queue)
		for i := 0; i < size; i++ {
			for j := 0; j < 4; j++ {
				x, y := queue[i][0]+dx[j], queue[i][1]+dy[j]
				if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != 1 {
					continue
				}
				grid[x][y] = 2
				queue = append(queue, [2]int{x, y})
			}
		}
		queue = queue[size:]
	}
	// 最后确定是否有遗漏的新鲜橘子没有被腐烂
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return res

}
