/*
454. 四数相加2
2023-1-10
link: https://leetcode.cn/problems/4sum-ii/
question:
	给你四个整数数组 nums1、nums2、nums3 和 nums4 ，数组长度都是 n ，
	请你计算有多少个元组 (i, j, k, l) 能满足：
		0 <= i, j, k, l < n
		nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0

answer:
	暴力解决：四层for循环，看那种符合要求,数组过长会超时。

	题目是四个独立的数组，只要找到nums1[i] + nums2[j] + nums3[k] + nums4[l] = 0就可以，
	不用考虑有重复的四个元素相加等于0的情况
	优化：
	1. 定义一个哈希map，key放nums1和nums2两数之和，value放nums1和nums2两数之和出现的次数
	2. 遍历nums1和nums2数组，统计两个数组元素之和，和出现的次数，放到map中。
	3. 定义int变量res，用来统计 nums1[i] + nums2[j] + nums3[k] + nums4[l] = 0 出现的次数。
	4. 在遍历nums3和nums4数组，找到如果 0-(nums3[k] + nums4[l]) 在map中出现过的话，
		就用res把map中key对应的value也就是出现次数统计出来。
	5. 最后返回统计值 res 就可以了
*/

// 暴力方法：37 / 132 个通过测试用例。若数组长度过长会超时。
// func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
// 	res := 0
// 	for i := 0; i < len(nums1); i++ {
// 		for j := 0; j < len(nums2); j++ {
// 			for k := 0; k < len(nums3); k++ {
// 				for l := 0; l < len(nums4); l++ {
// 					if nums1[i]+nums2[j]+nums3[k]+nums4[l] == 0 {
// 						res++
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return res
// }

// 优化：哈希思想
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 声明和初始化hash
	hash := map[int]int{}
	// 遍历前两个数组
	for _, val1 := range nums1 {
		for _, val2 := range nums2 {
			hash[val1+val2]++
		}
	}
	// 遍历后两个数组，统计数量
	res := 0
	for _, val3 := range nums3 {
		for _, val4 := range nums4 {
			res += hash[0-val3-val4]
		}
	}
	return res
}