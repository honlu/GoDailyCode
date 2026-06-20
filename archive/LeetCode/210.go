/*
210. 课程表2
2023-4-25 by lu
link: https://leetcode.cn/problems/course-schedule-ii/
question:
	现在你总共有 numCourses 门课需要选，记为 0 到 numCourses - 1。
	给你一个数组 prerequisites ，其中 prerequisites[i] = [ai, bi] ，
	表示在选修课程 ai 前 必须 先选修 bi 。

	例如，想要学习课程 0 ，你需要先完成课程 1 ，我们用一个匹配来表示：[0,1] 。
	返回你为了学完所有课程所安排的学习顺序。
	可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。

answer:
	第一遍做，没有思路。
	参考题解：尤其关于题解中拓扑节点保存的方式，然后深度遍历结合！

*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 不清楚思路
	// 拓扑排序，参考题解-深度优先搜索：栈-保留顺序，所有节点深度优先搜索
	edges := make([][]int, numCourses) // 边
	visited := make([]int, numCourses) // 边是否访问过
	result := []int{}                  // 栈
	valid := true
	var dfs func(u int)
	dfs = func(u int) { // 访问节点u
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0]) // 这个很重要，bi所对应的ai数组
	}

	fmt.Println(edges)

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	// 栈反转，即返回结果
	for i := 0; i < len(result)/2; i++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result
}