package greedy

import "fmt"

/*
3
376. 摆动序列
day:2022-6-19
update:: 2023-2-8
link:https://leetcode.cn/problems/wiggle-subsequence/
question:
	如果连续数字之间的差严格地在正数和负数之间交替，则数字序列称为摆动序列。
	第一个差（如果存在的话）可能是正数或负数。少于两个元素的序列也是摆动序列。

	给定一个整数序列，返回作为摆动序列的最长子序列的长度。
	通过从原始序列中删除一些（也可以不删除）元素来获得子序列，剩下的元素保持其原始顺序。
idea:
贪心的题目说简单有的时候就是常识，说难就难在都不知道该怎么用贪心。

本题大家如果要去模拟删除元素达到最长摆动子序列的过程，那指定绕里面去了，一时半会拔不出来。

而这道题目有什么技巧说一下子能想到贪心么？
其实也没有，类似的题目做过了就会想到。
此时大家就应该了解了：保持区间波动，只需要把单调区间上的元素移除就可以了
*/
// 局部符合，则整体可能满足
func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return 1
	}
	// 利用前一个差值和当前差值来实现摆动序列的长度
	// 注意当前差为0时，没有意义。
	var res, pre, cur int // 结果值，前一个差值，当前差值.初始均为0
	var path []int        // 一个摆动序列，测试使用
	res = 1
	path = append(path, nums[0])
	for i := 1; i < len(nums); i++ {
		cur = nums[i] - nums[i-1]
		if (cur > 0 && pre <= 0) || (cur < 0 && pre >= 0) {
			pre = cur // 注意这里，只在摆动变化的时候更新pre
			res++
			path = append(path, nums[i])
		}
	}
	fmt.Println(path) // 打印一个摆动序列。摆动序列长度没问题，但摆动序列可能有问题
	// 比如：例子[1,17,5,10,13,15,10,5,16,8]，输出[1 17 5 10 10 16 8]。10 10有问题
	// 暂时没解决，待思考
	return res
}

// 贪心算法解决摆动序列
func wiggleMaxLength(nums []int) int {
	var res, preDif, curDif int // 声明结果，前一个差值，当前差值
	res = 1                     // 为了避免最左面和最右面不好统计
	if len(nums) < 2 {
		return res
	}
	for i, j := 1, len(nums); i < j; i++ {
		curDif = nums[i] - nums[i-1]
		// 如果有正负则更新下标值即出现局部峰值
		if (curDif > 0 && preDif <= 0) || (curDif < 0 && preDif >= 0) {
			preDif = curDif
			res++
		}
	}
	return res
}
