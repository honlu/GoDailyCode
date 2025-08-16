package backtrack

/*
58-组合总和
https://leetcode.cn/problems/combination-sum/submissions/638147331/?envType=study-plan-v2&envId=top-100-liked
*/
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	var dfs func(target, index int)
	count := len(candidates)
	// 这样写法会有重复组合，如何避免重复组合呢？就是必须有一个开始下标，下标不能迁移
	dfs = func(target, index int) {
		if target == 0 {
			res = append(res, append([]int{}, path...)) // 注意这里细节
			return
		}
		for ; index < count; index++ {
			item := candidates[index]
			if target-item < 0 {
				continue
			}
			path = append(path, item)
			dfs(target-item, index)
			path = path[:len(path)-1]
		}
	}
	dfs(target, 0)
	return res
}
