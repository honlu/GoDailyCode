package backtrack77

func combine(n int, k int) [][]int {
	res := [][]int{}
	path := []int{}
	var backtrack func(int)
	backtrack = func(start int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(1)
	return res
}
