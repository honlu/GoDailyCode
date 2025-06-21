package hot100

/*
59-括号生成
https://leetcode.cn/problems/generate-parentheses/description/?envType=study-plan-v2&envId=top-100-liked
*/
func generateParenthesis(n int) []string {
	// 首先括号成对要先有左括号
	var res []string
	var path []rune
	var dfs func(left, right int)
	dfs = func(left, right int) {
		if left > n || right > n || right > left {
			return
		}
		if left+right == 2*n {
			res = append(res, string(path))
		}
		// 首先左括号进入
		path = append(path, '(')
		dfs(left+1, right)
		path = path[:len(path)-1]
		// 回溯之后，右括号进入
		path = append(path, ')')
		dfs(left, right+1)
		path = path[:len(path)-1]
	}
	dfs(0, 0)
	return res
}
