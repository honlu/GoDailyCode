/*
3、下一个更大元素
2022-10-13
link: 503-https://leetcode.cn/problems/next-greater-element-ii/
question:
	给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），
	输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，
	这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。
	如果不存在，则输出 -1。
answer:
	注意点：循环数组
	如何处理循环数组？ 遍历两次，直接通过下标求余法完成。
*/
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}
	stack := []int{}                   // 栈，存数组中元素的下标
	for i := 0; i < len(nums)*2; i++ { // 注意这里*2
		// 模拟遍历两边nums, 注意一下都是用i%len(nums)来操作
		for len(stack) > 0 && nums[i%len(nums)] > nums[stack[len(stack-1)]] {
			res[stack[len(stack)-1]] = nums[i%len(nums)] // 添加结果
			stack = stack[:len(stack)-1]                 // 出栈
		}
		stack = append(stack, i%len(nums))
	}
	return res
}