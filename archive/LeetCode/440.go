/*
440. 字典序的第k个小数字
2023-1-8
link: https://leetcode.cn/problems/k-th-smallest-in-lexicographical-order/
question:
	给定整数 n 和 k，返回  [1, n] 中字典序第 k 小的数字。
answer:
	注意：1 <= k <= n <= 10^9
	前缀树：深度优先遍历
	1作为根节点（前缀），子节点为10 - 19（以1为前缀）
	10作为根节点，子节点为100 - 109（以10为前缀）
	难。参考：https://leetcode.cn/problems/k-th-smallest-in-lexicographical-order/solution/pythonjavajavascriptgo-di-gui-by-himymbe-5mq5/
*/
// 暂时不会写的代码
func findKthNumber(n int, k int) int {
	// 小于等于n的以1开头的数有多少个?
	// 1 10-19 100-199 1000-1999 = 1111
	var dfs func(left, right int) int
	dfs = func(left, right int) int {
		if left > n {
			return 0
		}
		if right > n {
			right = n
		}
		return right - left + 1 + dfs(left*10, right*10+9)
	}

	cur := 1
	k -= 1
	for k > 0 {
		counts := dfs(cur, cur)
		// 当前节点中总数都小于需要的数，可以全部取走，bfs到同层下一点 (比如 1 -> 2)
		if counts <= k {
			k -= counts
			cur++
		} else { // 答案在当前节点的子节点中，取走当前根节点，dfs向下 (比如 1 -> 10)
			k--
			cur *= 10
		}
	}
	return cur
}

// 暴力生成字典序，试试通过几个用例：39 / 69
// func findKthNumber(n int, k int) int {
// 	ans := make([]int, 0, n) // 注意这种方式：n过大，会爆内存。
// 	for i, num := 0, 1; i < n; i++ {
// 		ans = append(ans, num)
// 		if num*10 <= n {
// 			num *= 10
// 		} else {
// 			for num%10 == 9 || num+1 > n {
// 				num /= 10
// 			}
// 			num++
// 		}
// 	}
// 	return ans[k]
// }