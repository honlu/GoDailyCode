package stack

/*
73-84.柱状图中最大的矩形
https://leetcode.cn/problems/largest-rectangle-in-histogram/description/?envType=study-plan-v2&envId=top-100-liked
*/
/*
思路：假设在某个柱子，如何判断以它为高度的右边界和左边界呢？
使用一个 单调递增栈（存放下标 index，不存高度）。
栈内保持 递增高度顺序，这样当出现更矮的柱子时，就能确定之前某个柱子的「右边界」。
*/

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0) // 添加一个哨兵，在末尾加一个高度为0的柱子，保证栈清空
	m := len(heights)
	var res int
	// 单调栈，递增
	var stack []int           // 栈中存放高度的位置
	stack = append(stack, -1) // 哨兵，方便计算左边界
	for i := 0; i < m; i++ {
		for len(stack) > 1 && heights[i] < heights[stack[len(stack)-1]] { // 当i的高度小于栈顶元素高度时，说明当前i是栈顶元素的矩阵右边界
			top := stack[len(stack)-1]   // 栈顶元素
			stack = stack[:len(stack)-1] // 出栈
			left := stack[len(stack)-1]  // 当前栈顶元素的左边界
			res = max(res, (i-1-left)*heights[top])
		}
		// 元素下标进栈
		stack = append(stack, i)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
