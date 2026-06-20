package p210

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, prerequisite := range prerequisites {
		next, prev := prerequisite[0], prerequisite[1]
		edges[prev] = append(edges[prev], next)
		indegree[next]++
	}

	queue := []int{}
	for course, degree := range indegree {
		if degree == 0 {
			queue = append(queue, course)
		}
	}

	order := make([]int, 0, numCourses)
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		order = append(order, course)
		for _, next := range edges[course] {
			indegree[next]--
			if indegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	if len(order) != numCourses {
		return []int{}
	}
	return order
}
