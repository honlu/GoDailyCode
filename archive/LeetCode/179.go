/*
179. 最大数
2022-12-13
link: https://leetcode.cn/problems/largest-number/
question:
	给定一组非负整数 nums，
	重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
answer:
	模拟排序过程。
	转为字符串，然后对字符串按照规则对比，然后排序。[关键]
	最后合并结果。
测试用例：
[10,2]
[3,30]
[3,30,34,5,9]
[3,30,34,5,9,3,30,34,5,9]
*/
package leetcode

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	// 转为字符串
	var sNums []string
	for _, val := range nums {
		sNums = append(sNums, strconv.Itoa(val))
	}
	// 排序
	sort.SliceStable(sNums, func(i, j int) bool {
		if sNums[i][0] > sNums[j][0] {
			return true
		} else {
			return sNums[i]+sNums[j] > sNums[j]+sNums[i] // 字符串拼接顺序后比较大小
		}
	})
	// 特殊情况，[0,0,0,0]
	if sNums[0] == "0" {
		return "0"
	}
	// 合并结果 strings.Join(sNums, "") 也可以
	var res string
	for _, s := range sNums {
		res += s
	}
	return res
}
