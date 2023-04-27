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
func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 排序,从小到大.方便排除重复元素
	var res [][]int
	temp := 0
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 { // 排序后三元组的第一个元素,一定不能大于0
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // 排除重复元素,和上一次枚举的数不相同
			continue
		}
		j, k := i+1, len(nums)-1 // 双指针
		for j < k {
			temp = nums[i] + nums[j] + nums[k]
			if temp < 0 { // 三数之和小于0,第一个值是要枚举的,最后一个值已经是最大的,所以只能改中间相对最小的
				j = j + 1 // 必须要有这个，才可以继续for循环，否则会死循环
				for j < k && nums[j] == nums[j-1] {
					j++
				}
			} else if temp > 0 { // 三数之和小于0,第一个值是要枚举的,中间已经是相对最小的,所以只能改最后最大的值
				k -= 1
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else { // 三数之和等于0,添加
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j += 1
				k -= 1
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return res
}