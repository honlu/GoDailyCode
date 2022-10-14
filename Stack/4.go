/*
4、接雨水
2022-10-13
link: 42-https://leetcode.cn/problems/trapping-rain-water/
question:
	给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，
	计算按此排列的柱子，下雨之后能接多少雨水。
answer:
	单调栈的思路不是很好理解！看原文！
	首先单调栈是按照行方向来计算雨水.
	其次使用单调栈内元素的顺序：
		选择从栈头（元素从栈头弹出）到栈底的顺序应该是从小到大的顺序。
*/

// 单调栈解法
func trap(height []int) int {
	if len(height) <= 2 { // 可以不写
		return 0
	}
	st := []int{0} // 利用切片模拟单调栈，存数组中元素的下标。计算的时候用下标对应柱子的高度
	res := 0
	for i := 1; i < len(height); i++ {
		if height[i] < height[st[len(st)-1]] { // 情况1
			st = append(st, i)
		} else if height[i] == height[st[len(st)-1]] { // 情况2
			st = st[:len(st)-1] // 这一句可以不加，效果一样，但处理相同的情况的思路却变了。
			st = append(st, i)
		} else { // 情况3 最关键！
			for len(st) > 0 && height[i] > height[st[len(st)-1]] {
				top := st[len(st)-1]
				st = st[:len(st)-1]
				if len(st) > 0 {
					h := height[i]
					if h > height[st[len(st)-1]] { // 这个比较关键！
						h = height[st[len(st)-1]]
					}
					h -= height[top]
					res += h * (i - st[len(st)-1] - 1)
				}
			}
			st = append(st, i)
		}
	}
	return res
}

