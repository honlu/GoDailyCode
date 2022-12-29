package greedy

/*
3、摆动序列
day:2022-6-19
link:https://leetcode.cn/problems/wiggle-subsequence/
idea:
贪心的题目说简单有的时候就是常识，说难就难在都不知道该怎么用贪心。

本题大家如果要去模拟删除元素达到最长摆动子序列的过程，那指定绕里面去了，一时半会拔不出来。

而这道题目有什么技巧说一下子能想到贪心么？
其实也没有，类似的题目做过了就会想到。
此时大家就应该了解了：保持区间波动，只需要把单调区间上的元素移除就可以了
*/

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
