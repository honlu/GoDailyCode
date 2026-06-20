package leetcode

/*
12-整数转罗马数字
https://leetcode.cn/problems/integer-to-roman/
*/
func intToRoman(num int) string {
	// 暴力枚举
	thousand := [4]string{"", "M", "MM", "MMM"}
	hundred := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	ten := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	unit := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	// 模拟
	return thousand[num/1000] + hundred[(num%1000)/100] + ten[(num%100)/10] + unit[num%10]
}
