package lcr_122

import "strings"

/*
题目：
  - https://leetcode.cn/problems/ti-huan-kong-ge-lcof/description/?envType=study-plan-v2&envId=coding-interviews
  - 假定一段路径记作字符串 path，其中以 "." 作为分隔符。
    现需将路径加密，加密方法为将 path 中的分隔符替换为空格 " "，请返回加密后的字符串。

示例 1：
输入：path = "a.aef.qerf.bb"
输出："a aef qerf bb"

解析：
- 本题考察字符串的替换操作，可以使用 strings.ReplaceAll() 方法进行替换。
*/
func pathEncrypting(path string) string {
	return strings.ReplaceAll(path, ".", " ")
}
