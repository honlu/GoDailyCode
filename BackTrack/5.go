package backtrack

/*
5、组合总和
day:2022-7-13
link: 39- https://leetcode.cn/problems/combination-sum/
question: 给定一个无重复元素的数组 candidates 和一个目标数 target ，
	找出 candidates 中所有可以使数字和为 target 的组合。
	candidates 中的数字可以无限制重复被选取。
answer:
	抽象成树型结构来解决。
	叶子节点返回条件：因为本题没有组合数量要求，仅仅是总和的限制，所以递归没有层数的限制，
		只要选取的元素总和超过target，就返回！
注意：下面代码为了逻辑清晰，使用了sum!也可以不用sum，使用target来做减法。
	这里startIndex控制循环的起始位置，也非常有趣！
	res二维切片使用指针也是一个很有趣的东西！不适用指针不行哦！
*/

// 在递归中传递下一个数字.与2.go代码不同
func combinationSum(candidates []int, target int) [][]int {
	var track []int
	var res [][]int
	backTrack(0, 0, target, candidates, track, &res)
	return res
}

// 递归函数
func backTrack(startIndex, sum, target int, candidates, track []int, res *[][]int) {
	// 递归的终止条件
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
		backTrack(i, sum, target, candidates, track, res) // 关键点，不用i+1了，标识可以重复读取当前的数字
		// 回溯
		track = track[:len(track)-1]
		sum -= candidates[i]
	}
}
