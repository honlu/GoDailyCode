/*
202.快乐数
2023-1-9
link: https://leetcode.cn/problems/happy-number/
question:
	编写一个算法来判断一个数 n 是不是快乐数。
	「快乐数」 定义为：
	对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
	然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
	如果这个过程 结果为 1，那么这个数就是快乐数。
	如果 n 是 快乐数 就返回 true ；不是，则返回 false 。

answer:
	利用哈希存储每次循环得到的结果。如果结果存在，就是存在了死循环不可能是快乐数。
*/
func isHappy(n int) bool {
	// 哈希表建立
	hash := map[int]bool{}
	temp := n
	for temp != 1 {
		temp = digitSum(temp) // 注意这里temp不能使用:=赋值，否则因为局部变量，for循环条件temp不变化
		if hash[temp] != true {
			hash[temp] = true
		} else {
			return false
		}
	}
	return true
}

// 一个数字各位值的平方和
func digitSum(num int) int {
	s := strconv.Itoa(num)
	res := 0
	for _, val := range s {
		res += int(val-'0') * int(val-'0')
	}
	return res
}