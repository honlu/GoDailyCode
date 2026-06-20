package hot100

/*
7-接雨水
双指针理解有一点抽象，需要好好想和看梳理关系
*/
func trap(height []int) int {
	// 从整体左右上来看
	left, right, leftMax, rightMax := 0, len(height)-1, 0, 0
	var ans int
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}
