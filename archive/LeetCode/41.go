/*
41. 缺失的第一个正数
2022-11-28
link: https://leetcode.cn/problems/first-missing-positive/
question:
	给你一个未排序的整数数组 nums ，
	请你找出其中没有出现的最小的正整数。
answer:
	注意时间复杂度有要求O(n),以及常熟级别额外空间。
*/
package main

import (
	"fmt"
	"sort"
)

func firstMissingPositive(nums []int) int {
	if len(nums) == 1 { // 特殊情况处理
		if nums[0] <= 0 || nums[0] > 1 {
			return 1
		} else {
			return 2
		}
	}
	// 长度大于2一般情况，首先排序
	sort.Ints(nums)
	temp := 0 // 然后有一个保存最小正整数的上一位数！
	for i := 0; i < len(nums); i++ {
		if nums[i] <= temp {
			continue
		} else if nums[i]-temp > 1 {
			break
		} else if nums[i]-temp == 1 {
			temp = nums[i]
		}
	}
	return temp + 1
}

func main() {
	nums := []int{7, 8, 9, 11, 12}
	res := firstMissingPositive(nums)
	fmt.Println(res)
}
