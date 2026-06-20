/*
402. 移掉K位数字
2022-12-20
link: https://leetcode.cn/problems/remove-k-digits/
question:
	给你一个以字符串表示的非负整数 num 和一个整数 k ，
	移除这个数中的 k 位数字，使得剩下的数字最小。
	请你以字符串形式返回这个最小的数字。
answer:
	虽然是中等题，但难度感觉到困难了。主要贪心和单调栈思想要融会贯通。
	参考：https://leetcode.cn/problems/remove-k-digits/solution/wei-tu-jie-dan-diao-zhan-dai-ma-jing-jian-402-yi-d/
	思路：
		从左至右扫描，当前扫描的数还不确定要不要删，入栈暂存。
		123531这样「高位递增」的数，肯定不会想删高位，会尽量删低位。
		432135这样「高位递减」的数，会想干掉高位，直接让高位变小，效果好。
		所以，如果当前遍历的数比栈顶大，符合递增，是满意的，让它入栈。
*/

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	// 栈声明和初始化
	var stack []rune
	for _, c := range num {
		// 当栈非空时，要判断是否高位递减，如果递减要元素出栈
		for len(stack) > 0 && k > 0 && stack[len(stack)-1] > c {
			stack = stack[:len(stack)-1]
			k--
		}
		if c == '0' && len(stack) == 0 { // 前导0不入栈
			continue
		}
		stack = append(stack, c) // 元素入栈
	}
	if k >= 0 && len(stack) > k {
		stack = stack[:len(stack)-k]
	} else if k >= 0 && len(stack) <= k {
		return "0"
	}
	return string(stack)
}