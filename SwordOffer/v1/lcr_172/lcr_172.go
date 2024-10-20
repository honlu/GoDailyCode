package lcr172

/*
LCR 172. 统计目标成绩的出现次数
某班级考试成绩按非严格递增顺序记录于整数数组 scores，请返回目标成绩 target 的出现次数。
*/

// FindTargetCount 统计目标成绩的出现次数
func FindTargetCount(scores []int, target int) int {
	res := 0
	left, right := 0, len(scores)
	mid := 0
	// 左闭右开写法,二分查找
	for left < right {
		mid = (left + right) / 2
		if scores[mid] < target {
			left = mid + 1
		} else if scores[mid] > target {
			right = mid
		} else {
			break
		}
	}
	if len(scores) == 0 || scores[mid] != target {
		return res
	}
	for index := mid; index < right && scores[index] == target; { // index < right要先判断，避免scores[index]越界，下同
		res++
		index++
	}
	for index := mid - 1; index >= left && scores[index] == target; {
		res++
		index--
	}
	return res
}

// FindTargetCount2 统计目标成绩的出现次数
func FindTargetCount2(scores []int, target int) int {
	count := 0
	for _, item := range scores {
		if item == target {
			count++
		} else if item > target {
			break
		}
	}
	return count
}
