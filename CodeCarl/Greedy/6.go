/*
6
55. 跳跃游戏
2022-10-12
update: 2023-2-11 by lu
link: https://leetcode.cn/problems/jump-game/
question:
	给定一个非负整数数组，你最初位于数组的第一个位置。
	数组中的每个元素代表你在该位置可以跳跃的最大长度。
	判断你是否能够到达最后一个位置。
answer:
	贪心算法局部最优解：每次取最大跳跃步数（取最大覆盖范围）
	整体最优解：最后得到整体最大覆盖范围，看是否能到终点。
	局部最优推出全局最优，找不出反例，试试贪心！
*/
// 贪心
// 记录每个索引可以到达的最大范围，每次遍历这个最大范围内是否有能到达最后下标的。不断更新索引以及最大范围
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxCover := nums[0] // 索引0可以达到的最大范围
	for i := 1; i <= maxCover; i++ {
		// 若当前索引到达最大范围比上次大，更新
		if i+nums[i] > maxCover {
			maxCover = i + nums[i]
		}
		// 判断是否可以到达最后一位
		if maxCover >= len(nums)-1 {
			return true
		}
	}
	return false
}