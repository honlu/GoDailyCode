/*
1004. 最大连续1的个数3
2023-4-5 by lu
link: https://leetcode.cn/problems/max-consecutive-ones-iii/
question:
	给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。
answer:
	方法一：参考题解使用前缀和、二分查找，这个思路很不熟悉呀！需要改进和提升
*/
func longestOnes(nums []int, k int) int {
	// 二分部查找和前缀和
	n := len(nums)
	temp := make([]int, n+1) // 保存前缀和的数组，temp[0]=0, temp[i+1] = temp[i]+(1-nums[i]),保存一个区间内0的个数，1-nums[i]就是把0变成1、1变成0
	res := 0
	for i, v := range nums {
		temp[i+1] = temp[i] + 1 - v
	}
	for right, v := range temp {
		left := sort.SearchInts(temp, v-k) // 这个很巧妙
		if right-left > res {
			res = right - left
		}
	}
	return res
}