package backtrack

/*
46-全排列
https://leetcode.cn/problems/permutations/description/?envType=study-plan-v2&envId=top-100-liked
*/
func permute(nums []int) [][]int {
	used := make(map[int]bool)
	var res [][]int
	var recursion func()
	var path []int
	recursion = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for _, num := range nums {
			if used[num] {
				continue
			}
			used[num] = true
			path = append(path, num)
			recursion()
			used[num] = false         // 回溯
			path = path[:len(path)-1] // 回溯
		}
	}
	recursion()
	return res
}
