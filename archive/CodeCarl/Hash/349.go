/*
349. 两个数组的交集
2023-1-8
link: https://leetcode.cn/problems/intersection-of-two-arrays/
question:
	给定两个数组 nums1 和 nums2 ，返回 它们的交集 。
	输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
answer:
	排序+哈希存储
*/
// 排序，哈希简单存储数字是否存在
// func intersection(nums1 []int, nums2 []int) []int {
// 	// 排序
// 	sort.Ints(nums1)
// 	sort.Ints(nums2)
// 	// 遍历数组1存储哈希
// 	hash := map[int]bool{}
// 	hash[nums1[0]] = true
// 	for i := 1; i < len(nums1); i++ {
// 		if nums1[i] == nums1[i-1] {
// 			continue
// 		}
// 		hash[nums1[i]] = true
// 	}
// 	// 遍历数组2保存交集
// 	res := []int{}
// 	if hash[nums2[0]] == true {
// 		res = append(res, nums2[0])
// 	}
// 	for i := 1; i < len(nums2); i++ {
// 		if nums2[i] == nums2[i-1] {
// 			continue
// 		}
// 		if hash[nums2[i]] == true {
// 			res = append(res, nums2[i])
// 		}
// 	}
// 	return res
// }

// 上面代码优化。不排序也可以
func intersection(nums1 []int, nums2 []int) []int {
	// 哈希表嫌存储nums1中的每个字符使用1计数，再判断nums2中字符在哈希表中是否存在，
	// 存在就保存到结果数组中，保存后使哈希变为0.避免结果多次存储
	var res []int
	hash := map[int]int{}
	for _, val := range nums1 {
		hash[val] = 1 // 这里很巧妙
	}
	for _, val := range nums2 {
		if hash[val] > 0 {
			res = append(res, val)
			hash[v]-- // 一个字符只添加一次
		}
	}
	return res
}