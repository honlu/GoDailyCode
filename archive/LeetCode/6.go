package leetcode

/*
6-Z字型变换
https://leetcode.cn/problems/zigzag-conversion/description/
*/
/*
思路：模拟操作
*/
// func convert(s string, numRows int) string {
// 	// 模拟操作
// 	var res [][]byte
// 	res = make([][]byte, numRows)
// 	for i := 0; i < len(s); {
// 		// 从上到下
// 		for j := 0; j < numRows && i < len(s); j++ {
// 			res[j] = append(res[j], s[i])
// 			i++
// 		}
// 		// 从左到右
// 		for j := numRows - 2; j > 0 && i < len(s); j-- {
// 			res[j] = append(res[j], s[i])
// 			i++
// 		}
// 	}
// 	var temp string
// 	for i := 0; i < numRows; i++ {
// 		temp += string(res[i])
// 	}
// 	return temp
// }

/*
找规律题：
1. 第一行和最后一行，都是2n-2的等差序列
2. 除去头尾两行的中间行是由两个2n-2等差序列混合组成
*/
func convert(s string, numRows int) string {
	var res []byte
	// 避免死循环,如果s="A",numsRows=1;s="AB",numsRows=1
	if numRows == 1 {
		return s
	}
	for i := 0; i < numRows; i++ {
		if i == 0 || i == numRows-1 { // 第一行或最后一行
			for j := i; j < len(s); j += 2*numRows - 2 { // 这里会死循环，如果s="A",numsRows=1
				res = append(res, s[j])
			}
		} else { // 中间行
			for j, k := i, 2*numRows-i-2; j < len(s) || k < len(s); j, k = j+2*numRows-2, k+2*numRows-2 {
				if j < len(s) {
					res = append(res, s[j])
				}
				if k < len(s) {
					res = append(res, s[k])
				}
			}
		}
	}
	return string(res)
}
