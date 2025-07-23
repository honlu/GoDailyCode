package leetcode

/*
17-电话号码的组合
https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
*/
/*
回溯
*/
func letterCombinations(digits string) []string {
	digitMap := map[rune][]rune{
		'2': []rune{'a', 'b', 'c'},
		'3': []rune{'d', 'e', 'f'},
		'4': []rune{'g', 'h', 'i'},
		'5': []rune{'j', 'k', 'l'},
		'6': []rune{'m', 'n', 'o'},
		'7': []rune{'p', 'q', 'r', 's'},
		'8': []rune{'t', 'u', 'v'},
		'9': []rune{'w', 'x', 'y', 'z'},
	}
	var res []string
	if digits == "" {
		return res
	}
	var path []rune
	var dfs func(index int)
	digitRune := []rune(digits)
	dfs = func(index int) {
		if len(path) == len(digitRune) || index > len(digitRune) {
			res = append(res, string(path))
			return
		}
		for _, item := range digitMap[digitRune[index]] {
			path = append(path, item)
			dfs(index + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
