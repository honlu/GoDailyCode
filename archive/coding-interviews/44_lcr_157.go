package main

// LCR 157. 套餐内商品的排列顺序
func goodsOrder(goods string) []string {
	seen := make(map[string]bool)
	var dfs func(remaining, current string)
	// 回溯函数
	dfs = func(remaining, current string) {
		// 如果剩余字符串长度为1，则将当前字符串加入结果集
		if len(remaining) == 1 {
			res := current + remaining
			if !seen[res] {
				seen[res] = true
			}
		}
		// 遍历剩余字符串
		for i := 0; i < len(remaining); i++ {
			// 将当前字符加入当前字符串
			temp2 := current + string(remaining[i])
			// 删除当前字符
			temp := remaining[:i] + remaining[i+1:] // 删除当前字符
			dfs(temp, temp2)
		}
	}
	dfs(goods, "")
	var res []string
	for item := range seen {
		res = append(res, item)
	}
	return res
}
