/*
6、跳跃游戏
2022-10-12
link: 55-https://leetcode.cn/problems/jump-game/
question:
	给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个位置。
answer:
	贪心算法局部最优解：每次取最大跳跃步数（取最大覆盖范围）
	整体最优解：最后得到整体最大覆盖范围，看是否能到终点。
	局部最优推出全局最优，找不出反例，试试贪心！
*/

func canJump(nums []int) bool {
	cover := 0
	if len(nums) == 1 {
		return true // 只有一个元素，一定能达到
	}
	for i := 0; i <= cover; i++ { // 注意这里是小于等于
		if i+nums[i] > cover {
			cover = i + nums[i]
		}
		if cover >= len(nums)-1 { // 关键：说明已经覆盖到终点
			return true
		}
	}
	return false
}