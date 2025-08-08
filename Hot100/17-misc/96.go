package misc

/*
96-136.只出现一次的数字
https://leetcode.cn/problems/single-number/?envType=study-plan-v2&envId=top-100-liked
*/
/*
Go：排序+遍历，时间复杂度O（nlogn）
*/
// func singleNumber(nums []int) int {
// 	sort.Ints(nums)
// 	var res int = nums[len(nums)-1]
// 	for i := 0; i < len(nums)-1; {
// 		if nums[i] != nums[i+1] {
// 			res = nums[i]
// 			break
// 		}
// 		i += 2
// 	}
// 	return res
// }

/*
哈希：时间复杂度O(n)
*/
// func singleNumber(nums []int) int {
// 	numMap := make(map[int]bool)
// 	for i := 0; i < len(nums); i++ {
// 		if numMap[nums[i]] {
// 			delete(numMap, nums[i])
// 		} else {
// 			numMap[nums[i]] = true
// 		}
// 	}
// 	var res int
// 	for item, _ := range numMap {
// 		res = item
// 	}
// 	return res
// }
func singleNumber(nums []int) int {
	// 异或运算：0与任何数异或，结果都是任何数；两个相同的数据异或为0
	var res int = nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
