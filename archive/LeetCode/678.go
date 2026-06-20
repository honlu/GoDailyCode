/*
678. 有效的括号字符串
2023-4-3 by lu
link: https://leetcode.cn/problems/valid-parenthesis-string/
question:
	给定一个只包含三种字符的字符串：（ ，） 和 *，写一个函数来检验这个字符串是否为有效字符串。有效字符串具有如下规则：

		任何左括号 ( 必须有相应的右括号 )。
		任何右括号 ) 必须有相应的左括号 ( 。
		左括号 ( 必须在对应的右括号之前 )。
		* 可以被视为单个右括号 ) ，或单个左括号 ( ，或一个空字符串。
		一个空字符串也被视为有效字符串。

answer:
	栈思想解决：本题关键多了一个星号的特性。
	所以需要两个栈：左括号栈、星号栈。
	从左到右遍历字符串：
		如果遇到左括号，则将当前下标存入左括号栈。
		如果遇到星号，则将当前下标存入星号栈。
		如果遇到右括号，则需要有一个左括号或星号和右括号匹配，由于星号也可以看成右括号或者空字符串，
		因此当前的右括号应优先和左括号匹配，没有左括号时和星号匹配：
			如果左括号栈不为空，则从左括号栈弹出栈顶元素；
			如果左括号栈为空且星号栈不为空，则从星号栈弹出栈顶元素；
			如果左括号栈和星号栈都为空，则没有字符可以和当前的右括号匹配，返回false。
	遍历结束，左括号栈和星号栈可能还有多余，要再次判断，是否星号可以匹配多余的左括号。
*/
func checkValidString(s string) bool {
	leftStack, starStack := []int{}, []int{} // 两个栈：分别存左括号、星号对应下标
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			leftStack = append(leftStack, i)
		} else if s[i] == '*' {
			starStack = append(starStack, i)
		} else { // 为右括号
			if len(leftStack) > 0 {
				leftStack = leftStack[:len(leftStack)-1] // 优先：左括号栈出元素
			} else if len(starStack) > 0 {
				starStack = starStack[:len(starStack)-1] // 型号栈出元素
			} else {
				return false
			}
		}
	}
	// 如果左括号还有多余的，需要星号来匹配
	i, j := len(leftStack)-1, len(starStack)-1
	if i > j {
		return false
	}
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if leftStack[i] > starStack[j] {
			return false
		}
	}
	return true
}