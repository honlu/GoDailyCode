package backtrack

/*
5、组合总和
day:2022-7-13
还没来得及注释和运行测试
*/

// 在递归中传递下一个数字.与2.go代码不同
func combinationSum(candidates []int, target int) [][]int {
	var track []int
	var res [][]int
	backtracking(0, 0, target, candidates, track, &res)
	return res
}

func backtracking(startIndex, sum, target int, candidates, track []int, res *[][]int) {
	// 终止条件
	if sum == target {
		temp := make([]int, len(track))
		copy(temp, track)         // 深拷贝
		*res = append(*res, temp) // 放入结果集
		return
	}
	if sum > target {
		return
	}
	// 回溯
	for i := startIndex; i < len(candidates); i++ {
		// 更新路径集合和sum
		track = append(track, candidates[i])
		sum += candidates[i]
		// 递归
		backtracking(i, sum, target, candidates, track, res)
		// 回溯
		track = track[:len(track)-1]
		sum -= candidates[i]
	}
}
