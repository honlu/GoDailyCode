package greed

/*
78-跳跃游戏
*/
func canJump(nums []int) bool {
	// 能接触到的索引下标的最大跳跃覆盖范围
	if len(nums) == 1 {
		return true
	}
	index := 0
	maxCoverr := index + nums[index]
	for i := index + 1; i <= maxCoverr; i++ {
		if i+nums[i] > maxCoverr {
			maxCoverr = i + nums[i]
		}
		if maxCoverr >= len(nums)-1 {
			return true
		}
	}
	return false
}
