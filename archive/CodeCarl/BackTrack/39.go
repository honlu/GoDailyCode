package backtrack

/*
5
39. 组合总和
day:2022-7-13
update: 2023-2-3 by lu
link: https://leetcode.cn/problems/combination-sum/
question:
	给定一个无重复元素的数组 candidates 和一个目标数 target ，
	找出 candidates 中所有可以使数字和为 target 的组合。
	candidates 中的数字可以无限制重复被选取。
answer:
	注意本题：没有数量要求，可以无限重复，但总和有要求，所以简接的也有个数的限制。

	抽象成树型结构来解决。
	叶子节点返回条件：因为本题没有组合数量要求，仅仅是总和的限制，所以递归没有层数的限制，
		只要选取的元素总和超过target，就返回！
注意：下面代码为了逻辑清晰，使用了sum!也可以不用sum，使用target来做减法。
	这里startIndex控制循环的起始位置，也非常有趣！
	res二维切片使用指针也是一个很有趣的东西！不适用指针不行哦！
*/

// 回溯，建议不加path参数，内存消耗更少！这里为了方便理解，回溯函数参数加了path参数。
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int // 符合结果集
	var path []int  // 符合结果
	// 回溯函数定义和创建:
	/*
		参数定义：
			targetSum:题目要求的目标值
			sum：当前结果路径的和
			start: 下一层for循环搜索的起始位置 [关键理解]]
			path: 可加可不加
	*/
	var backTrack func(targetSum, sum, start int, path []int)
	backTrack = func(targetSum, sum, start int, path []int) {
		// base case
		if sum > targetSum { // 剪枝操作，当然也可以加到下面的for循环条件里。当然是要整个candidate数组有序的情况下。
			return
		}
		if sum == targetSum {
			res = append(res, append([]int{}, path...))
			return
		}
		// 回溯算法标准步骤
		for i := start; i < len(candidates); i++ {
			sum += candidates[i]               // 处理
			path = append(path, candidates[i]) // 处理
			backTrack(targetSum, sum, i, path) //递归。不用i+1,表示可以重复读取当前的数[精华]
			sum -= candidates[i]               // 回溯
			path = path[:len(path)-1]          // 回溯
		}
	}
	backTrack(target, 0, 0, path)

	return res
}

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
