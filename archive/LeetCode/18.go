package leetcode

import "sort"

/*
18-四数之和
https://leetcode.cn/problems/4sum/description/
*/
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 去重
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] { // 去重，注意这里是j>i+1,不是0
				continue
			}
			for k, l := j+1, len(nums)-1; k < l; { // 双指针
				if k > j+1 && nums[k] == nums[k-1] { // 去重
					k++
					continue
				}
				temp := nums[i] + nums[j] + nums[k] + nums[l]
				if temp == target {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
					k++ // 继续
					l--
				} else if temp < target {
					k++
				} else {
					l--
				}
			}
		}
	}
	return res
}
