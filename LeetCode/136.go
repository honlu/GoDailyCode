/*
136.只出现一次的数字
2022-12-20
link: https://leetcode.cn/problems/single-number/
question:
	给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，
	其余每个元素均出现两次。找出那个只出现了一次的元素。
	要求：线性时间复杂度，常量额外空间
answer:
	先排序，然后遍历解决
*/
package leetcode

import "sort"

func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums) // 排序
	i := 0
	for i < len(nums)-1 { // 判断
		if nums[i+1] == nums[i] {
			i += 2
		} else {
			return nums[i]
		}
	}
	return nums[len(nums)-1] // 在最后一位
}
