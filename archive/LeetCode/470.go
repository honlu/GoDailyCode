/*
470. 用Rand7()实现Rand10()
2022-12-5
link: https://leetcode.cn/problems/implement-rand10-using-rand7/
question:
	给定方法 rand7 可生成 [1,7] 范围内的均匀随机整数，
	试写一个方法 rand10 生成 [1,10] 范围内的均匀随机整数。
	你只能调用 rand7() 且不能调用其他方法。
	请不要使用系统的 Math.random() 方法。

answer:
	https://leetcode.cn/problems/implement-rand10-using-rand7/solution/xiang-xi-si-lu-ji-you-hua-si-lu-fen-xi-zhu-xing-ji/
	题解参考！
	重点理解思路+ 概率计算
*/

// 最简单方式，首先一个数产生范围再1-x范围内，x要大于10.然后通过判断，再次生成符号要求的。
// 不过这里简单但效率不高
func rand10() int {
	num := (rand7()-1)*7 + rand7()
	for num > 10 {
		num = (rand7()-1)*7 + rand7()
	}
	return num
}