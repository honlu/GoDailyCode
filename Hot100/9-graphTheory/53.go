package hot100

/*
53-课程表
https://leetcode.cn/problems/course-schedule/?envType=study-plan-v2&envId=top-100-liked
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 判断图判断是否有环，prerequisites就是边
	// 每个入度为0的节点都要判断
	graph := make(map[int][]int)
	d := make(map[int]int) // 每个节点有多少入度
	for _, pre := range prerequisites {
		a, b := pre[1], pre[0]
		graph[a] = append(graph[a], b)
		d[b]++ // 注意是b的入度++
	}
	var queue []int
	// 入度为0的节点入队
	for i := 0; i < numCourses; i++ {
		if d[i] == 0 {
			queue = append(queue, i)
		}
	}
	// bfs广搜
	var cnt int
	for len(queue) > 0 {
		size := len(queue)
		cnt += size
		for i := 0; i < size; i++ {
			for _, j := range graph[queue[i]] {
				d[j]--
				if d[j] == 0 {
					queue = append(queue, j)
				}
			}
		}
		queue = queue[size:]
	}
	return cnt == numCourses
}
