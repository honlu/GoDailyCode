/*
7. 整数反转
2023-3-4 by lu
link：https://leetcode.cn/problems/reverse-integer/
question:
	给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
	如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
	假设环境不允许存储 64 位整数（有符号或无符号）。
answer:
	注意最大最小32位数字的表示；
	求余、数字相乘相加更新；

*/
func reverse(x int) int {
	res := 0
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 { // 这个判断条件很精妙：只需要判断是否最小值小于min/10,最大值是否大于max/10,因为后面需要乘以10来更新
			return 0
		}
		digit := x % 10
		res = res*10 + digit
		x /= 10
	}
	return res
}