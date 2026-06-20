package lcr_138

import "strings"

/*
题目：有效数字
- https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/description/?envType=study-plan-v2&envId=coding-interviews
- 有效数字（按顺序）可以分成以下几个部分：
 1. 若干空格
 2. 一个 小数 或者 整数
 3. （可选）一个 'e' 或 'E' ，后面跟着一个 整数
 4. 若干空格
    小数（按顺序）可以分成以下几个部分：
 1. （可选）一个符号字符（'+' 或 '-'）
 2. 下述格式之一：
 1. 至少一位数字，后面跟着一个点 '.'
 2. 至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
 3. 一个点 '.' ，后面跟着至少一位数字
    整数（按顺序）可以分成以下几个部分：
 1. （可选）一个符号字符（'+' 或 '-'）
 2. 至少一位数字
    部分有效数字列举如下：["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]
    部分无效数字列举如下：["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]

- 给你一个字符串 s ，如果 s 是一个 有效数字 ，请返回 true 。

  - 示例 1：
    输入：s = "0"
    输出：true
  - 示例 2：
    输入：s = "e"
    输出：false
  - 示例 3：
    输入：s = "."
    输出：false

- 解析：
*/
func validNumber(s string) bool {
	// 去除前后可能存在空格
	s = strings.TrimSpace(s)
	// 是否存在数字
	number := false
	// 是否存在小数点
	dot := false
	// 是否存在e
	e := false
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			number = true
		} else if s[i] == '.' {
			if dot || e {
				return false
			}
			dot = true
		} else if s[i] == 'e' || s[i] == 'E' {
			if e || !number {
				return false
			}
			e = true
			number = false
		} else if s[i] == '+' || s[i] == '-' {
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else {
			return false
		}
	}
	return number
}

// 解释上面的代码
// 1. 去除前后可能存在空格
// 2. 是否存在数字
// 3. 是否存在小数点
// 4. 是否存在e
// 5. 遍历字符串
// 6. 如果是数字，则number为true
// 7. 如果是小数点，则判断是否已经存在小数点或者e，如果存在则返回false，否则dot为true
// 8. 如果是e，则判断是否已经存在e或者没有数字，如果存在则返回false，否则e为true，number为false
// 9. 如果是+或者-，则判断是否是第一个字符或者前一个字符是e或者E，如果不是则返回false
// 10. 其他情况返回false
// 11. 如果遍历结束，number为true，则返回true，否则返回false
// 详细一句话，总结上面的内容：
// 1. 去除前后可能存在空格，初始化number、dot、e为false，遍历字符串，如果是数字则number为true，如果是小数点则判断是否已经存在小数点或者e，如果存在则返回false，否则dot为true，如果是e则判断是否已经存在e或者没有数字，如果存在则返回false，否则e为true，number为false，如果是+或者-，则判断是否是第一个字符或者前一个字符是e或者E，如果不是则返回false，其他情况返回false，如果遍历结束，number为true，则返回true，否则返回false，即返回是否存在数字，如果存在数字则返回true，否则返回false。
// 2. 时间复杂度O(n)，空间复杂度O(1)。
// 本题只需要遍历一次字符串s，所以时间复杂度为O(n)。
// 本题只需要使用常数个变量，所以空间复杂度为O(1)。
// 所以总的时间复杂度为O(n)，空间复杂度为O(1)。
