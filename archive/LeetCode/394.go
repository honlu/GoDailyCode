//  中等题-394字符串解码
func decodeString(s string) string {
	// 使用递归实现
	i := 0
	var getDigits func() int // 获取数字的函数
	getDigits = func() int {
		ret := 0 // return 简写
		for ; s[i] >= '0' && s[i] <= '9'; i++ {
			ret = ret*10 + int(s[i]-'0')
		}
		return ret
	}
	var getString func() string
	getString = func() string {
		if i == len(s) || s[i] == ']' { // 递归结束条件
			return ""
		}
		cur := s[i]
		times := 1 // 数字或者说重复次数
		ret := ""
		if cur >= '0' && cur <= '9' {
			times = getDigits()
			i++                              // 此时i为'['索引，所以++移到下一位
			str := getString()               // 递归
			i++                              // 跳出循环时i为']'索引，所以++移到下一位
			ret = strings.Repeat(str, times) // 利用了包函数
		} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {
			ret = string(cur)
			i++
		}
		return ret + getString() // 这个有些深套
	}
	res := getString()
	return res
}