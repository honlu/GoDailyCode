package dynamicprogramming

/*
80-70：爬楼梯
https://leetcode.cn/problems/climbing-stairs/description/?envType=study-plan-v2&envId=top-100-liked
*/
func climbStairs(n int) int {
	// 规律公式：f(n) = f(n-1) + f(n-2)
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1
	for i := 2; i < n+1; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}
