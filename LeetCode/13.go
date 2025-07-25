package leetcode

/*
13-罗马数字转整数
https://leetcode.cn/problems/roman-to-integer/description/
*/
/*
思路：模拟
*/
func romanToInt(s string) int {
	hash := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// // 字符串从后往前遍历，组成个位、十位、百位、千位等
	// var res int
	// var last int
	// for i := len(s) - 1; i >= 0; i-- {
	// 	temp := hash[s[i]]
	// 	if temp < last { // 若左边元素比右边小，则减
	// 		res -= temp
	// 	} else { // 若左边元素小于等于右边，则加
	// 		res += temp
	// 	}
	// 	last = temp
	// }

	// 从前后遍历，后面比当前大就减，比当前小或等于就+
	var res int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && hash[s[i]] < hash[s[i+1]] {
			res -= hash[s[i]]
		} else {
			res += hash[s[i]]
		}
	}
	return res
}
