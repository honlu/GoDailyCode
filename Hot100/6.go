package hot100

import "sort"

/*
6-三数之和
第一遍没有做出来
*/
func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-2; i++ {
		item := nums[i]
		if item > 0 {
			break
		}
		if i > 0 && item == nums[i-1] { // 这里要处理：跳过重复的第一个数，避免结果中出现重复的三元组。
			continue
		}

		res = append(res, TwoSum(nums[i+1:], -item)...)
	}
	return res
}

func TwoSum(temp []int, target int) [][]int {
	// temp从小到大，双指针找temp中两数之和为target的
	left, right := 0, len(temp)-1
	var res [][]int
	for left < right {
		a, b := temp[left], temp[right]
		if a+b == target {
			res = append(res, []int{a, b, -target})
			// 这里要处理:1. 避免for一直循环，内存切漏.2.跳过重复[重点]
			for left < right && temp[left] == a {
				left++
			}
			for left < right && temp[right] == b {
				right--
			}
		} else if a+b > target {
			right--
		} else {
			left++
		}
	}
	return res
}
