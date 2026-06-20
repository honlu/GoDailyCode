package greed

/*
79:45-跳跃游戏2-最小跳跃数
https://leetcode.cn/problems/jump-game-ii/?envType=study-plan-v2&envId=top-100-liked
*/
/*
BFS（广度优先搜索）分层遍历：
把每个下标当作图中的一个节点，每次从当前位置跳到 i+1 ~ i+nums[i] 这些位置，类似从一个节点出发能连出去多个边。
*/
func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0 // 已经在终点了，不需要跳
	}

	steps := 0         // 当前跳跃次数
	start, end := 0, 0 // 当前 BFS 层的左右边界（闭区间）

	for end < n-1 {
		steps++         // 进入下一跳层
		furthest := end // 当前层可跳的最远距离

		// 遍历当前层的所有位置 i
		for i := start; i <= end; i++ {
			if i+nums[i] > furthest {
				furthest = i + nums[i] // 更新下一层的边界
			}
		}

		// 下一层的 start 和 end
		start = end + 1
		end = furthest
	}
	return steps
}
