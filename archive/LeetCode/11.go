package leetcode

/*
11. 盛最多水的容器
2022-12-25
link: https://leetcode.cn/problems/container-with-most-water/
question:
	给定一个长度为 n 的整数数组 height 。
	有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
	找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
	返回容器可以储存的最大水量。
answer:
	贪心和双指针！
	为了获取最大值，采用双指针。
	i，j对应高度较小的那个先向内移动，不断计算面积，更新最大面积
*/
// func maxArea(height []int) int {
// 	var i, j int = 0, len(height) - 1
// 	var res int = 0
// 	for i < j {
// 		small = min(height[i], height[j])
// 		res = min(res, small*(j-i))
// 		if small == height[i] {
// 			i++
// 		} else {
// 			j--
// 		}
// 	}
// 	return res
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

/*
方程计算两点储水最大: max((j-i)*min(height[i],height[j]))
双指针
贪心优化
*/
func maxArea(height []int) int {
	// 贪心、双指针
	i, j := 0, len(height)-1
	var res int
	for i < j {
		res = max(res, (j-i)*min(height[i], height[j]))
		if height[i] < height[j] { // 贪心逻辑
			i++
		} else {
			j--
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
