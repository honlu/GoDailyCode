/*
1. 两数之和
2023-1-9
link: https://leetcode.cn/problems/two-sum/
question:
	给定一个整数数组nums和一个整数目标值target，请你在该数组中找出和为目标值target的那两个整数，
	并返回它们的数组下标。

	你可以假设每种输入只会对应一个答案。
	但是，数组中同一个元素在答案里不能重复出现。

	你可以按任意顺序返回答案。

answer:
	注意：假设每种输入只对应一个答案。两个整数。
	所以使用哈希，key为数组中的值，value为对应数组索引。
*/
func twoSum(nums []int, target int) []int {
	// 哈希表建立。由于value类型是int，默认为0，很容易和数组索引0的冲突，所以value存储索引+1，方便判断。
	hash := map[int]int{} // value存储数组索引+1
	for i := 0; i < len(nums); i++ {
		if hash[target-nums[i]] != 0 {
			return []int{i, hash[target-nums[i]] - 1}
		} else {
			hash[nums[i]] = i + 1
		}
	}
	return []int{}
}

// 代码优化：go支持多返回值，map[key]可以返回两个元素，第二元素ok可以判断是否存储。
// 这样就不需要value是索引下标+1，可以直接是索引下标
// func twoSum(nums []int, target int) []int {
// 	// 哈希表建立。由于value类型是int，默认为0，很容易和数组索引0的冲突，所以value存储索引+1，方便判断。
// 	hash := map[int]int{} // value存储数组索引+1
// 	for i := 0; i < len(nums); i++ {
// 		if index, ok := hash[target-nums[i]]; ok {
// 			return []int{i, index}
// 		} else {
// 			hash[nums[i]] = i
// 		}
// 	}
// 	return []int{}
// }