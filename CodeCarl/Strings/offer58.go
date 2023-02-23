/*
剑指offer58. 左旋转字符串
2023-1-11
link: https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
question:
	字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。
	请定义一个函数实现字符串左旋转操作的功能。
	比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。

answer:
	暴力：直接切片切分+拼接.
	空间复杂度：O(n)

	难度提升：不能申请额外空间，只能在本串上操作。
*/
// 暴力
// func reverseLeftWords(s string, n int) string {
// 	return s[n:] + s[:n]
// }

// 代码优化：通过局部反转+整体反转 达到左旋转的目的
/*
具体步骤为：
	反转区间为前n的子串
	反转区间为n到末尾的子串
	反转整个字符串
这个时间复杂度也是O(n)呀！
*/
func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	reverseString(b[:n])
	reverseString(b[n:])
	reverseString(b)
	return string(b)
}
func reverseString(s []byte) {
	// 双指针
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
