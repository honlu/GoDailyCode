package backtrack

/*
61-分割回文串
https://leetcode.cn/problems/palindrome-partitioning/description/?envType=study-plan-v2&envId=top-100-liked
*/
func partition(s string) [][]string {
	// 将s转为byte数组，依次回溯每个开始节点到某个结束节点是否是回文串
	bs := []byte(s)
	var res [][]string
	n := len(bs)
	var path []string
	var dfs func(start, end int)
	dfs = func(start, end int) {
		if start >= n {
			res = append(res, append([]string{}, path...)) // 注意这里细节
			return
		}
		for i := end; i <= n; i++ {
			temp := bs[start:i]
			if !isHW(temp) {
				continue
			}
			path = append(path, string(temp))
			dfs(i, i+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 1)
	return res
}

func isHW(target []byte) bool {
	n := len(target)
	if n == 0 {
		return false
	}
	for i := 0; i < n/2; i++ {
		if target[i] != target[n-1-i] { // 注意细节
			return false
		}
	}

	return true
}
