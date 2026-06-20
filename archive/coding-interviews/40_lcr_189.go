package main

// LCR 189. 机械累加器
func mechanicalAccumulator(target int) int {
	// 递归 代替循环
	var ans int
	var dfs func(n int) bool
	dfs = func(n int) bool {
		ans += n
		return n > 0 && dfs(n-1) //通过go的判断左项优先提前退出法则
	}
	dfs(target)
	return ans
}
