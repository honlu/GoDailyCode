package backtrack

/*
4
17. 电话号码的字母组合
day:2022-6-19
update: 2023-2-1
link:https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
question:
	给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
	给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
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
var digitsMap = map[int]string{ // 创建map，保存数字和字母的对应
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

// 回溯
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res := []string{} // 结果集
	var path []byte   // 符合结果, go中string是不可变的
	// 回溯函数声明
	var backTrack func(start int, path []byte)
	backTrack = func(start int, path []byte) {
		// base case
		if start == len(digits) {
			res = append(res, string(path))
			return
		}
		// 单层逻辑
		digit := int(digits[start] - '0') // 将start指向的digits里的数字转为int
		letters := digitsMap[digit]       // 去数字对应的字符集
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i]) // 处理
			backTrack(start+1, path)        // 递归，注意start+1, 下一层要处理的数字
			path = path[:len(path)-1]       // 回溯
		}
	}

	backTrack(0, path)
	return res
}

// 把回溯藏在递归的参数里,这种写法，很不直观,不建议写
func letterCombinations(digits string) []string {
	n := len(digits)
	if n <= 0 || n > 4 { // 题目提示要求,小心这里等于0的条件
		return nil
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
