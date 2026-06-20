/*
15. 三数之和
2023-1-10
link: https://leetcode.cn/problems/3sum/
question:
	给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
	同时还满足 nums[i] + nums[j] + nums[k] == 0 。请
	你返回所有和为 0 且不重复的三元组。
	注意：答案中不可以包含重复的三元组。

answer:
	注意：这是在一个数组内寻找，没有要求索引，而是要求元素.
	两层for循环就可以确定 a 和b 的数值了，可以使用哈希法来确定 0-(a+b) 是否在 数组里出现过，其实这个思路是正确的，
	但是我们有一个非常棘手的问题，就是题目中说的不可以包含重复的三元组。
	把符合条件的三元组放进vector中，然后再去重，这样是非常费时的，
	很容易超时，也是这道题目通过率如此之低的根源所在。

	所以可以先排序,然后使用双指针,来寻找.难点:三元组去重,细节比较难拿捏,很容易出bug.
*/
func threeSum(nums []int) [][]int {
	// 先排序
	sort.Ints(nums)
	// 哈希判断其中一个值是否存在
	var res [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 { // 排序之后如果第一个元素已经大于零，那么不可能凑成三元组
			break
		}
		if i > 0 && nums[i] == nums[i-1] { // 三元组元素a去重
			continue
		}
		hash := map[int]bool{}
		for j := i + 1; j < len(nums); j++ { // 注意从i+1开始
			if j > i+2 && nums[j] == nums[j-1] && nums[j-1] == nums[j-2] { // 注意这里细节j与j-1比较,j-1与j-2比较
				continue // 三元组元素b去重
			}
			c := 0 - nums[i] - nums[j]
			if hash[c] == true {
				res = append(res, []int{nums[i], nums[j], c})
				hash[c] = false
			} else {
				hash[nums[j]] = true
			}
		}
	}
	return res
}