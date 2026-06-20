/*
18. 四数之和
2023-1-10
link: https://leetcode.cn/problems/4sum/
question:
	给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。
	请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b], nums[c], nums[d]] （若两个四元组元素一一对应，则认为两个四元组重复）：
	0 <= a, b, c, d < n
	a、b、c 和 d 互不相同
	nums[a] + nums[b] + nums[c] + nums[d] == target
	你可以按 任意顺序 返回答案 。

answer:
	双指针,将O(n^4)转为O(n^3)
*/
// 双指针法
func fourSum(nums []int, target int) [][]int {
	// 排序
	sort.Ints(nums)
	//
	var res [][]int
	for i := 0; i < len(nums)-3; {
		for j := i + 1; j < len(nums)-2; {
			// 双指针
			left, right := j+1, len(nums)-1
			for left < right {
				if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left = sort.SearchInts(nums, nums[left]+1) // 第三个数去重,即找比nums[left]大的下一位
				}
			}
			// 同理更新第二个数,去重
			j = sort.SearchInts(nums, nums[j]+1)
			// SearchInts在已排序的整型数切片中搜索x，并返回Search指定的索引。
			// 返回值是插入x的索引，如果x不存在,就会返回比x大的数的index(可以是len(a))。
		}
		// 同理更新第一个数,去重
		i = sort.SearchInts(nums, nums[i]+1)
	}
	return res
}