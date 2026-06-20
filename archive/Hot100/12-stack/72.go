package stack

/*
739-每日温度
*/
// func dailyTemperatures(temperatures []int) []int {
// 	// 暂时不知道怎么用栈解决，暴力试试，新的测试case过不去。
// 	answer := make([]int, len(temperatures))
// 	for i := 0; i < len(temperatures)-1; i++ {
// 		for j := i + 1; j < len(temperatures); j++ {
// 			if temperatures[j] > temperatures[i] {
// 				answer[i] = j - i
// 				break
// 			}
// 		}
// 	}
// 	return answer
// }

func dailyTemperatures(temperatures []int) []int {
	// 单调栈，占中存元素的下标索引「占中存什么，怎么进栈、出栈是关键」
	stack := []int{0}
	count := len(temperatures)
	answer := make([]int, count) // 默认为0，最后元素没出栈的就是0
	for i := 1; i < count; i++ {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] { // 遇到比它的元素就出栈
			answer[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return answer
}
