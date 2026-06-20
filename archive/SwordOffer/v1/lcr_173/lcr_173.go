package lcr173

/*
题目：https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/description/?envType=study-plan-v2&envId=coding-interviews
*/
func takeAttendance(records []int) int {
	// 找一个位置和值不相同的位置
	// 二分：闭区间方法
	left := 0
	right := len(records) - 1
	for left <= right { // 注意等号
		mid := (left + right) / 2
		if records[mid] == mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
