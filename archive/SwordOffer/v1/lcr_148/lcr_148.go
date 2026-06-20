package lcr148

/*
题目：验证图书取出顺序
- https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/description/?envType=study-plan-v2&envId=coding-interviews
- 现在图书馆有一堆图书需要放入书架，并且图书馆的书架是一种特殊的数据结构，只能按照 一定 的顺序 放入 和 拿取 书籍。
给定一个表示图书放入顺序的整数序列 putIn，请判断序列 takeOut 是否为按照正确的顺序拿取书籍的操作序列。你可以假设放入书架的所有书籍编号都不相同。

- 标签：栈
- 解题思路：模拟，验证出栈顺序是否正确
*/
func validateBookSequences(putIn []int, takeOut []int) bool {
	length := len(putIn)
	if length == 0 {
		return true
	}
	var stack []int
	i := 0
	for _, item := range putIn {
		stack = append(stack, item)                               // 模拟进栈
		for len(stack) > 0 && stack[len(stack)-1] == takeOut[i] { // 模拟出栈
			stack = stack[:len(stack)-1]
			i++
		}
	}
	return len(stack) == 0 // 判断是否全部出栈
}
