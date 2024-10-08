package lcr146

/*
题目：螺旋遍历二维数组
- https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/
- 给定一个二维数组 array，请返回「螺旋遍历」该数组的结果。

螺旋遍历：从左上角开始，按照 向右、向下、向左、向上 的顺序 依次 提取元素，然后再进入内部一层重复相同的步骤，直到提取完所有元素
- 标签：数组
- 解题思路：模拟
*/
func spiralArray(array [][]int) []int {
	var res []int
	if len(array) == 0 {
		return res
	}
	top, down := 0, len(array)-1
	left, right := 0, len(array[0])-1
	for top <= down && left <= right { // 注意等于号
		for i := left; i <= right; i++ {
			res = append(res, array[top][i])
		}
		top++
		for i := top; i <= down; i++ {
			res = append(res, array[i][right])
		}
		right--
		// 这里必须加，特殊情况处理.避免非方阵
		if top > down || left > right {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, array[down][i])
		}
		down--
		for i := down; i >= top; i-- {
			res = append(res, array[i][left])
		}
		left++
	}

	return res
}
