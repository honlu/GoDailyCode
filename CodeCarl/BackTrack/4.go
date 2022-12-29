package backtrack

/*
4、电话号码的字母组合
day:2022-6-19
link:https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
idea:
从示例上来说，输入"23"，最直接的想法就是两层for循环遍历了吧，正好把组合的情况都输出了。
如果输入"233"呢，那么就三层for循环，如果"2333"呢，就四层for循环.......

遇到可以利用for循环的层数写出来，回溯就可以登场了！

理解本题后，要解决如下三个问题：
1、数字和字母如何映射
	[使用map或这二维数组]
2、两个字母就两个for循环，三个字符我就三个for循环，以此类推，然后发现代码根本写不出来
	[回溯法解决n个for循环的问题，回溯三部曲]
3、输入1 * #按键等等异常情况
	[排除这些异常情况]
*/

// 回溯解决字母组合
func letterCombinations(digits string) []string {
	n := len(digits)
	if n <= 0 || n > 4 { // 题目提示要求,小心这里等于0的条件
		return nil
	}
	digitsMap := map[string]string{ // 创建map，保存数字和字母的对应
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	var res []string                              // 声明结果变量
	var backtracking func(index int, temp string) // 声明回溯函数，三部曲：1参数 2、终止条件 3、遍历
	backtracking = func(index int, temp string) {
		if index == n { // 终止条件
			res = append(res, temp)
		} else {
			digit := string(digits[index])
			letters := digitsMap[digit]
			for i, m := 0, len(letters); i < m; i++ { // 遍历和回溯
				backtracking(index+1, temp+string(letters[i]))
			}
		}
	}
	backtracking(0, "")
	return res
}
