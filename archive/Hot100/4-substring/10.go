package hot100

/*
10-和为k的子数组
前缀和+哈希表
*/
func subarraySum(nums []int, k int) int {
	// 注意子数组是数字中元素的连续非空序列，就不能排序以及考虑滑动窗口。
	// 字节一面，腾讯一面题
	// 需要理解前缀和，整理「前缀和」模板吗？
	// 如何保存前缀和的结果和索引的关系呢？
	ans := 0
	sumMap := make(map[int]int) // key:和，value:可能数
	sumMap[0] = 1               // 和为0，默认为1种
	sum := 0
	for _, num := range nums {
		sum += num
		ans += sumMap[sum-k]
		sumMap[sum]++
	}
	return ans
}
