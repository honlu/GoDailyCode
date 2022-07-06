package backtrack

/*
3. 组合总和III
day:2022-6-16
link:https://leetcode.cn/problems/combination-sum-iii/
idea:在[1,2,3,4,5,6,7,8,9]这个集合中找到和为n的k个数的组合。
本题k相当于了树的深度，9（因为整个集合就是9个数）就是树的宽度。
*/

// 组合枚举-递归法
func combinationSum3(k int, n int) [][]int {
	var ans [][]int
	var temp []int
	var dfs func(cur, rest int)
	dfs = func(cur, rest int) {
		// 找到一个答案
		if len(temp) == k && rest == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		// 剪枝：跳过的数字过多，后面已经无法选到k个数字
		if len(temp)+10-cur < k || rest < 0 {
			return
		}
		// 跳过当前数字
		dfs(cur+1, rest)
		// 选当前数字
		temp = append(temp, cur)
		dfs(cur+1, rest-cur)
		temp = temp[:len(temp)-1]
	}
	dfs(1, n)
	return ans
}
