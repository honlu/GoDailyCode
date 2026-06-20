/*
150. 逆波兰表达式求值
2023-1-13
link: https://leetcode.cn/problems/evaluate-reverse-polish-notation/
question:
	给你一个字符串数组tokens，表示一个根据逆波兰表示法表示的算术表达式。
	请你计算该表达式。返回一个表示表达式值的整数。

	注意：
		有效的算符为 '+'、'-'、'*' 和 '/' 。
		每个操作数（运算对象）都可以是一个整数或者另一个表达式。
		两个整数之间的除法总是 向零截断 。
		表达式中不含除零运算。
		输入是一个根据逆波兰表示法表示的算术表达式。
		答案及所有中间计算结果可以用 32 位 整数表示。

answer:
	关键要理解：逆波兰表示法。以及和栈的关系。
	逆波兰表达式：是一种后缀表达式，所谓后缀就是指运算符写在后面。
	平常使用的算式则是一种中缀表达式，如 ( 1 + 2 ) * ( 3 + 4 ) 。
	该算式的逆波兰表达式写法为 ( ( 1 2 + ) ( 3 4 + ) * ) 。
	逆波兰表达式主要有以下两个优点：
		去掉括号后表达式无歧义，上式即便写成 1 2 + 3 4 + * 也可以依据次序计算出正确结果。
		适合用栈操作运算：遇到数字则入栈；遇到运算符则取出栈顶两个数字进行计算，并将结果压入栈中。
*/
func evalRPN(tokens []string) int {
	// 栈实现：符号不进栈，遇到符号栈头部两个元素，分别是符号左右的运算项。
	hash := map[string]bool{"+": true, "-": true, "*": true, "/": true}
	stack := make([]int, 0)
	for _, val := range tokens {
		if hash[val] == true {
			left, right := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch val {
			case "+":
				stack = append(stack, left+right)
			case "-":
				stack = append(stack, left-right)
			case "*":
				stack = append(stack, left*right)
			case "/":
				stack = append(stack, left/right)
			}
		} else {
			temp, _ := strconv.Atoi(val)
			stack = append(stack, temp)
		}
	}
	// tokens若符合要求，栈最后只剩结果值
	return stack[0]
}

// 上面代码还可以优化！