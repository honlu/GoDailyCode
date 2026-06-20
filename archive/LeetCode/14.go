/*
14. 最长公共前缀
2022-12-7
link: https://leetcode.cn/problems/longest-common-prefix/
question:
	编写一个函数来查找字符串数组中的最长公共前缀。
	如果不存在公共前缀，返回空字符串 ""。
answer:
	遍历
*/
// func longestCommonPrefix(strs []string) string {
// 	if len(strs) == 1 { // 特殊情况 ["a"]
// 		return strs[0]
// 	}
// 	var res string
// 	temp := 0
// 	for i := 0; i < len(strs[0]); i++ {
// 		for j := 1; j < len(strs); j++ {
// 			if i < len(strs[j]) && strs[j][i] == strs[0][i] {
// 				if j == len(strs)-1 { // 特殊情况 ["flower","flower","flower","flower"]
// 					temp = i + 1
// 				}
// 				continue
// 			} else {
// 				temp = i
// 				goto Loop
// 			}
// 		}
// 	}
// Loop:
// 	res = strs[0][:temp]
// 	return res
// }

/*
以第一个元素为主，遍历
*/
func longestCommonPrefix(strs []string) string {
	m := len(strs)
	var res []byte
	for i := 0; i < len(strs[0]); i++ {
		temp := strs[0][i]
		for j := 1; j < m; j++ {
			if i >= len(strs[j]) || strs[j][i] != temp {
				return string(res)
			}
		}
		res = append(res, temp)
	}
	return string(res)
}