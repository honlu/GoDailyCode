package leetcode

import (
	"math"
	"sort"
)

/*
16. 最接近的三数之和
2023-3-24
link: https://leetcode.cn/problems/3sum-closest/
question:
	给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
	返回这三个数的和。
	假定每组输入只存在恰好一个解
answer:
	n个数选择三个数可以是随机的，但我们遍历需要有技巧性，让其可以即按顺序又有逻辑。
	所以先排序，然后一次选择一个数首先数字，再通过双指针在其后面区域选择两个数字，是他们和最接近目标值。
*/
// func threeSumClosest(nums []int, target int) int {
// 	sort.Ints(nums)
// 	delta := math.MaxInt               // 目标值与三树之和的偏差
// 	for i := 0; i < len(nums)-2; i++ { // 注意减2
// 		sum := nums[i] + twoSumClosest(nums, i+1, target-nums[i])
// 		if abs(sum-target) < abs(delta) {
// 			delta = sum - target
// 		}
// 	}
// 	return target + delta
// }

// func twoSumClosest(nums []int, start int, target int) int {
// 	delta := math.MaxInt
// 	l, r := start, len(nums)-1
// 	for l < r {
// 		sum := nums[l] + nums[r]
// 		if abs(sum-target) < abs(delta) {
// 			delta = sum - target
// 		}
// 		if sum > target {
// 			r--
// 		} else { // 这里小心bug，不能单独写if sum <target
// 			l++
// 		}
// 	}
// 	return target + delta
// }

// func abs(num int) int {
// 	if num < 0 {
// 		return -num
// 	}
// 	return num
// }

/*
排序+双指针
*/
func threeSumClosest(nums []int, target int) int {
	// 一样的，先排序，然后固定一个元素，再双指针
	res, diff := 0, math.MaxInt
	m := len(nums)
	sort.Ints(nums)
	for i := 0; i < m-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, m-1; j < k; {
			if abs(nums[i]+nums[j]+nums[k]-target) < diff {
				res = nums[i] + nums[j] + nums[k]
				diff = abs(nums[i] + nums[j] + nums[k] - target)
			}
			if nums[i]+nums[j]+nums[k] == target {
				return nums[i] + nums[j] + nums[k]
			} else if nums[i]+nums[j]+nums[k] < target {
				j++
			} else {
				k--
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
