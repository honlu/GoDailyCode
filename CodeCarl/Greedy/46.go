/*
7
45. 跳跃游戏2
2022-10-12
2023-2-12 update by lu
link: https://leetcode.cn/problems/jump-game-ii/
question:
	给定一个非负整数数组，你最初位于数组的第一个位置。
	数组中的每个元素代表你在该位置可以跳跃的最大长度。
	你的目标是使用最少的跳跃次数到达数组的最后一个位置。
answer:
	贪心的思路
	局部最优：当前可移动距离尽可能多走，如果还没到终点，步数再加一。
	整体最优：一步尽可能多走，从而达到最小步数。
	注意：要从覆盖范围出发，不管怎么跳，覆盖范围内一定是可以跳到的，
	以最小的步数增加覆盖范围，覆盖范围一旦覆盖了终点，得到的就是最小步数！
	需要统计两个覆盖范围，当前这一步的最大覆盖和下一步最大覆盖。
*/

/*
方法一：思路
移动下标达到了当前覆盖的最远距离下标时，步数就要加一，来增加覆盖距离。
最后的步数就是最少步数。
注意，特殊情况。当移动下标达到了当前覆盖的最远距离下标时：
如果当前覆盖最远距离下标不是是集合终点，步数就加一，还需要继续走。
如果当前覆盖最远距离下标就是是集合终点，步数不用加一，因为不能再往后走了。
*/
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	cur := 0  // 当前覆盖最远距离下标
	next := 0 // 下一步覆盖最远距离下标
	res := 0  // 记录走的最大步数
	for i := 0; i < len(nums); i++ {
		if nums[i]+i > next {
			next = nums[i] + i // 更新下一步覆盖最远距离下标
		}
		if i == cur { // 遇到当前覆盖最远距离下标
			if cur < len(nums)-1 { // 如果当前覆盖最远距离下标不是终点
				res++      // 需要走下一步
				cur = next // 更新当前覆盖最远距离下标（相当于加油了）
				if next >= len(nums)-1 {
					break // 下一步的覆盖范围已经可以达到终点，结束循环
				}
			} else { // 当前覆盖最远距到达集合终点，不用做res++操作了，直接结束
				break
			}
		}
	}
	return res
}

/*
代码优化
*/
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	res := 0
	cover, nextCover := 0, 0
	for i := 0; i < len(nums)-1; i++ { // 精髓在于控制移动下标i只移动到len(nums)-2的位置
		if nums[i]+i > nextCover {
			nextCover = nums[i] + i
		}
		if i == cover {
			cover = nextCover
			res++
		}
	}
	return res
}