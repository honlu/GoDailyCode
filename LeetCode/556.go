/*
556. 下一个更大元素2
2023-5-8 by lu
link: https://leetcode.cn/problems/next-greater-element-iii/
question:
	给你一个正整数 n ，请你找出符合条件的最小整数，其由重新排列 n 中存在的每位数字组成，并且其值大于 n 。
	如果不存在这样的正整数，则返回 -1 。
	注意 ，返回的整数应当是一个 32 位整数 ，如果存在满足题意的答案，
	但不是 32 位整数 ，同样返回 -1 。
answer:
	当数字是从高位到地位是递减的，返回-1.
	当存在符合条件的，要从地位开始往高位遍历，当发现x与后面的数字不是递增的时，
	就要从后面数字中找到第一个比x大的值，互换。(这种思路不对，比如123654，正确结果是124356，错误为124653 )
	因为还缺少一个步骤，将x位之后的数字反转！
	（还要注意小于号，还是小于等于号的bug）
	[一下代码自己写的，没有像题解一样简化。需要优化]
*/
package leetcode

import (
	"fmt"
	"math"
)

func nextGreaterElement(n int) int {
	// 双指针
	var res int
	temp := []int{} // 保存n从地位到高位的数
	for n > 0 {
		temp = append(temp, n%10)
		n /= 10
	}
	fmt.Println(temp)
	i := 1
	for ; i < len(temp); i++ {
		if temp[i] >= temp[i-1] { // 注意等于号
			continue
		} else {
			j := 0
			for temp[j] <= temp[i] { // 寻找temp[j],注意等于号。 12443322
				j++
			}
			temp[j], temp[i] = temp[i], temp[j] // 互换i,j指向的数
			// 反转[0,i-1]的数
			for k, l := 0, i-1; k < l; k, l = k+1, l-1 {
				temp[k], temp[l] = temp[l], temp[k]
			}
			fmt.Println(temp)
			break
		}
	}
	if i == len(temp) {
		return -1
	} else {
		for j := len(temp) - 1; j >= 0; j-- {
			res = res*10 + temp[j]
			// 注意32位整数，res不能大于2^31-1
			if res > math.MaxInt32 {
				return -1
			}
		}
	}
	return res
}
