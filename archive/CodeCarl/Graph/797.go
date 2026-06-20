/*
题目：797. 所有可能的路径 - 力扣（LeetCode）
https://leetcode.cn/problems/all-paths-from-source-to-target/description/
要求：给你一个有 n 个节点的 有向无环图（DAG），请你找出所有从节点 0 到节点 n-1 的路径并输出（不要求按特定顺序）
	graph[i] 是一个从节点 i 可以访问的所有节点的列表（即从节点 i 到节点 graph[i][j]存在一条有向边）。

题解：使用深度搜索在邻接矩阵存储的表结构中探索起点到终点的路径。
*/
func allPathsSourceTarget(graph [][]int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(cur, end int)
	dfs = func(cur, end int) {
		// 结束条件
		if cur == end {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp) // 注意这里添加元素
			return
		}
		// 循环
		for i := 0; i < len(graph[cur]); i++ {
			path = append(path, graph[cur][i])
			dfs(graph[cur][i], end)   // 递归
			path = path[:len(path)-1] // 回溯
		}
	}
	// 从0起点开始
	path = append(path, 0)
	dfs(0, len(graph)-1)
	return res
}