package leetcode

/*
6-Z字型变换
https://leetcode.cn/problems/zigzag-conversion/description/
*/
func convert(s string, numRows int) string {
	// 模拟操作
	var res [][]byte
	res = make([][]byte, numRows)
	for i := 0; i < len(s); {
		// 从上到下
		for j := 0; j < numRows && i < len(s); j++ {
			res[j] = append(res[j], s[i])
			i++
		}
		// 从左到右
		for j := numRows - 2; j > 0 && i < len(s); j-- {
			res[j] = append(res[j], s[i])
			i++
		}
	}
	var temp string
	for i := 0; i < numRows; i++ {
		temp += string(res[i])
	}
	return temp
}
