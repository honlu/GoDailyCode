package hot100

/*
5-盛更多的水
双指针和贪心
*/
func maxArea(height []int) int {
	// 双指针,贪心只移动小的一边
	left, right := 0, len(height)-1
	var res int
	for left < right {
		small := min(height[left], height[right])
		res = max(res, small*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}
