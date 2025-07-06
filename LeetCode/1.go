package leetcode

/*
1-两数之和
https://leetcode.cn/problems/two-sum/description/
方法思路：
1. 暴力
2. 排序+双指针
3. 通过map「哈希表」记录存在的数。时间复杂度：O(n)
*/
func twoSum(nums []int, target int) []int {
	// 首先遍历map存储value,index
	valueIndex := make(map[int]int)
	for i, item := range nums {
		if index, ok := valueIndex[target-item]; ok {
			return []int{i, index}
		}
		valueIndex[item] = i //放在if之后，避免使用重复
	}
	return []int{}
}
