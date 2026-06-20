package lcr181

import "strings"

/*
	题目：字符串中的单词反转

- https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof/
- 你在与一位习惯从右往左阅读的朋友发消息，他发出的文字顺序都与正常相反但单词内容正确，为了和他顺利交流你决定写一个转换程序，把他所发的消息 message 转换为正常语序。

注意：输入字符串 message 中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

- 限制：0 <= s 的长度 <= 10000
- 注意：本题与主站 151 题相同：https://leetcode-cn.com/problems/reverse-words-in-a-string/
*/
func reverseMessage(message string) string {
	res := strings.Split(message, " ")
	var temp []string
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] != "" {
			temp = append(temp, res[i])
		}
	}
	return strings.Join(temp, " ")
}
