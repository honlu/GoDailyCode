/*
剑指offer5. 替换空格
2023-1-10
link: https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
question:
	请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
answer:
	%20算是一个byte或者rune吗？不算，那就不可以直接替换。算是一个字符串或者说byte数组
*/
// 套用strings函数
func replaceSpace(s string) string {
	// 使用函数，通过空格切分
	st := strings.Split(s, " ") // 返回一个字符串切片
	res := st[0]
	for i := 1; i < len(st); i++ { // 然后拼接
		res += "%20" + st[i]
	}
	return res
}