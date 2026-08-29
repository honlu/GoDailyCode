package lcr162

// 暴力会超时，待优化
/*
class Solution:
    def countDigitOne(self, n: int) -> int:
        num = 0
        for i in range(1,n+1):
            tmp = str(i)
            num += tmp.count('1') # count函数，统计一个字符串中某个字符出现次数
        return num
*/

// 数学：典型的 按位统计 / 数位规律题
func digitOneInNumber(num int) int {
	/*
		统计 1 ~ num 中数字 1 出现的总次数。
		核心思想：
		不要枚举每一个数字，而是分别统计：
			个位出现多少次 1
			十位出现多少次 1
			百位出现多少次 1
			千位出现多少次 1
			...
		digit 表示当前正在统计的数位：
			digit = 1    → 个位
			digit = 10   → 十位
			digit = 100  → 百位
			...
		对于当前数位，把 num 拆成：
			high | cur | low
		例如：
			num = 2304
			统计百位时 digit = 100
			high = 23
			cur  = 0
			low  = 4
			即：
			23 | 0 | 04
			     ↑
			    百位
	*/
	res := 0
	// 从个位开始逐位统计
	for digit := 1; digit <= num; digit *= 10 {
		// 当前位右边的数字
		//
		// 例如：
		// num = 2314
		// digit = 100
		//
		// low = 14
		low := num % digit
		// 当前位数字
		//
		// 2314 / 100 = 23
		// 23 % 10 = 3
		//
		// 所以 cur = 3
		cur := (num / digit) % 10
		// 当前位左边的数字
		//
		// 2314 / 1000 = 2
		//
		// 所以 high = 2
		high := num / (digit * 10)
		/*
			接下来根据 cur 和 1 的大小关系，
			分成三种情况。

			为什么一定是三种？

				cur < 1
				cur == 1
				cur > 1

			当前位能出现多少次 1，
			关键就取决于当前位已经走到了哪里。
		*/
		if cur == 0 {
			/*
				情况一：cur < 1
				当前位还没有走到 1。
				当前位出现 1 的完整轮数只有 high 轮，
				每一轮持续 digit 次。
				所以：
					count = high * digit
			*/
			res += high * digit
		} else if cur == 1 {
			/*
				情况二：cur == 1
				前面的 high 个完整周期贡献：
					high * digit
				当前这一轮已经走到了 1，
				还能贡献：
					low + 1
				因此：
					count = high*digit + low + 1
			*/
			res += high*digit + low + 1
		} else {
			/*
				情况三：cur > 1
				说明当前这一轮的数字 1 已经完整经过。
				所以除了 high 个完整周期，
				还要额外增加一个完整周期：
					count = (high + 1) * digit
			*/
			res += (high + 1) * digit
		}
	}
	return res
}
