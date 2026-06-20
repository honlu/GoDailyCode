/*
5、柱状图中最大的矩形
2022-10-14
link: 84-https://leetcode.cn/problems/largest-rectangle-in-histogram/
question:
	给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
	求在该柱状图中，能够勾勒出来的矩形的最大面积。
answer:
	和“接雨水”题目原理上相似，但细节有所差异。
	42. 接雨水 是找每个柱子左右两边第一个大于该柱子高度的柱子，
	而本题是找每个柱子左右两边第一个小于该柱子的柱子

	接雨水的单调栈从栈头（元素从栈头弹出）到栈底的顺序应该是从小到大的顺序。

	本题是要找每个柱子左右两边第一个小于该柱子的柱子，所以从栈头（元素从栈头弹出）
	到栈底的顺序应该是从大到小的顺序！才能保证栈顶元素找到左右两边第一个小于栈顶元素的柱子。
*/

func largestRectangleArea(heights []int) int {
	// 数组头部和尾部添加0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	// 栈初始化，保存数组中元素的下标
	stack := []int{0}
	res := 0
	for i := 1; i < len(heights); i++ {
		if heights[i] > heights[stack[len(stack)-1]] { // 情况1
			stack = append(stack, i) // 进栈
		} else if heights[i] == heights[stack[len(stack)-1]] { // 情况2
			stack = stack[:len(stack)-1] // 出栈。可以不加，效果一样，思路不同
			stack = append(stack, i)
		} else { // 情况3
			for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
				top := stack[len(stack)-1]   // 临时保存栈顶元素
				stack = stack[:len(stack)-1] // 出栈
				left := stack[len(stack)-1]  // 新的栈顶元素
				right := i
				if res < (right-left-1)*heights[top] { // 更新结果.注意这里right-left-1是减一
					res = (right - left - 1) * heights[top]
				}
			}
			stack = append(stack, i) // 进栈
		}
	}
	return res
}