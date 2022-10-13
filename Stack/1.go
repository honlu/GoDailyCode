/*
1、每日温度
2022-10-13
link:739-https://leetcode.cn/problems/daily-temperatures/
question:
	请根据每日气温列表，重新生成一个列表。对应位置的输出为：
		要想观测到更高的气温，至少需要等待的天数。
	如果气温在这之后都不会升高，请在该位置用 0 来代替。
answer:
	单调栈解法：单调栈的本质是空间换时间，
		因为在遍历的过程中需要用一个栈来记录右边第一个比当前元素高的元素，
		优点是只需要遍历一次。

	适用场景：通常是一维数组，要寻找任一个元素的右边或者
		左边第一个比自己大或者小的元素的位置，此时我们就要想到可以用单调栈了。

	单调栈：新的元素比栈顶元素小就都要进展，若大就要从栈顶依次将小的栈内元素出栈
*/

// 单调栈解法
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	// 单调栈,初始化栈顶元素为第一个下标索引0
	stack := []int{0}
	for i := 1; i < len(temperatures); i++ {
		if temperatures[i] < temperatures[stack[len(stack)-1]] { // 情况1：当前遍历元素小于栈顶元素
			stack = append(stack, i) // 进栈
		} else if temperatures[i] == temperatures[stack[len(stack)-1]] { // 情况2：当前遍历元素等于栈顶元素
			stack = append(stack, i) // 进栈
		} else { // 情况3：当前遍历元素大于栈顶元素.
			for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
				res[stack[len(stack)-1]] = i - stack[len(stack)-1]
				stack = stack[:len(stack)-1] // 出栈
			}
			stack = append(stack, i) // 进栈
		}
	}
	return res
}

// 暴力解法，两层for循环。Go可以通过
// func dailyTemperatures(temperatures []int) []int {
// 	lens := len(temperatures)
// 	answer := make([]int, len(temperatures))
// 	for i := 0; i < lens-1; i++ { // 最后一个默认为0
// 		for j := i + 1; j < lens; j++ {
// 			if temperatures[j] > temperatures[i] {
// 				answer[i] = j - i
// 				break // 跳出j的循环
// 			}
// 		}
// 	}
// 	return answer
// }