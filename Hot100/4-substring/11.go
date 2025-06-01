package hot100

/*
11-滑动窗口最大的值
双端队列，滑动窗口
*/

func maxSlidingWindow(nums []int, k int) []int {
	// 如何快速获取滑动窗口内的最大值
	// 方法二：双端队列，保存nums的索引
	// 队列底部保存最大值的索引，索引头部进行新元素或旧元素的出
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	windows := make([]int, 0, k) // 保存nums的索引
	result := make([]int, 0, len(nums)-k)
	for i, item := range nums {
		if i >= k && windows[0] <= i-k { // 注意边界问题=, 窗口左侧出
			windows = windows[1:]
		}
		for len(windows) > 0 && nums[windows[len(windows)-1]] < item { // 队列位置的元素较小，出队.注意元素获取：nums[windows[len(windows)-1]]
			windows = windows[:len(windows)-1]
		}
		windows = append(windows, i) // 窗口右侧进，当前元素的索引进队
		if i+1 >= k {                // 窗口右移一次，加入窗口内最大值
			result = append(result, nums[windows[0]])
		}
	}
	return result
}
