package leetcode

import "sort"

/*
15. 三数之和
2022-11-11
2023-4-27 updated, labeled as star by lu
link: https://leetcode.cn/problems/3sum/
question:
	给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]]
	满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。
	请你返回所有和为 0 且不重复的三元组。
answer:
	注意：不可以包含重复的三元组.
	排序+双指针
	注意：双指针使用技巧，如何遍历，这里可能是个困惑点！
*/
// func threeSum(nums []int) [][]int {
// 	sort.Ints(nums) // 排序,从小到大.方便排除重复元素
// 	var res [][]int
// 	temp := 0
// 	for i := 0; i < len(nums)-2; i++ {
// 		if nums[i] > 0 { // 排序后三元组的第一个元素,一定不能大于0
// 			break
// 		}
// 		if i > 0 && nums[i] == nums[i-1] { // 排除重复元素,和上一次枚举的数不相同
// 			continue
// 		}
// 		j, k := i+1, len(nums)-1 // 双指针
// 		for j < k {
// 			temp = nums[i] + nums[j] + nums[k]
// 			if temp < 0 { // 三数之和小于0,第一个值是要枚举的,最后一个值已经是最大的,所以只能改中间相对最小的
// 				j = j + 1 // 必须要有这个，才可以继续for循环，否则会死循环
// 				for j < k && nums[j] == nums[j-1] {
// 					j++
// 				}
// 			} else if temp > 0 { // 三数之和小于0,第一个值是要枚举的,中间已经是相对最小的,所以只能改最后最大的值
// 				k -= 1
// 				for j < k && nums[k] == nums[k+1] {
// 					k--
// 				}
// 			} else { // 三数之和等于0,添加
// 				res = append(res, []int{nums[i], nums[j], nums[k]})
// 				j += 1
// 				k -= 1
// 				for j < k && nums[j] == nums[j-1] {
// 					j++
// 				}
// 				for j < k && nums[k] == nums[k+1] {
// 					k--
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

func threeSum(nums []int) [][]int {
	// 输出的顺序和三元组的顺序并不重要，所以先对数字进行排序
	sort.Ints(nums)
	// 采用双指针。先固定最小的一个i位置，然后利用双指针在右边找到所有不重复的j,k
	var res [][]int
	m := len(nums)
	for i := 0; i < m-2; i++ {
		if nums[i] > 0 { // i位置是最小的，若nums[i]>0,则三数之和必大于0
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // i去重
			continue
		}

		j, k := i+1, m-1
		for j < k {
			if nums[i]+nums[j]+nums[k] == 0 { // 关键逻辑
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j+1] == nums[j] { // 去重
					j++ // 右移动
				}
				for j < k && nums[k-1] == nums[k] {
					k-- // 左移动
				}
				j++ // 注意上面是利用nums[j+1]=nums[j]，j++，所以执行完之后这里还要+1
				k-- //
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				j++
			}
		}
	}

	return res
}
