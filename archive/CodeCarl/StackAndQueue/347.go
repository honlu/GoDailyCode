/*
347. 前k个高频元素
2023-1-13
link: https://leetcode.cn/problems/top-k-frequent-elements/
question:
	给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。
	你可以按 任意顺序 返回答案。
answer:
	哈希表存储，然后再排序。
	Go中map是无序，所以思路：对key排序，再遍历key输出value。
	需要额外空间保存排序的key.
*/
func topKFrequent(nums []int, k int) []int {
	// 哈希
	hash := map[int]int{}
	key := []int{}
	// 遍历原数组,初始化hash,以及添加key
	for _, val := range nums {
		if hash[val] == 0 {
			key = append(key, val)
		}
		hash[val]++
	}
	// 然后对key进行排序。关键在这里，排序是根据哈希对应的值大小来对key切片排序
	// 注意高级排序要使用sort.Slice，才有第二个排序条件函数参数。sort.Ints无
	sort.Slice(key, func(i, j int) bool {
		return hash[key[i]] > hash[key[j]] //从大到下排序，这里关键
	})
	// 保存结果
	return key[:k]
}