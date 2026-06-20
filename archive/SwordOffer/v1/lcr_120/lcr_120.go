package lcr120

/*
题目：寻找文件副本
- https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/description/?envType=study-plan-v2&envId=coding-interviews
设备中存有 n 个文件，文件 id 记于数组 documents。若文件 id 相同，则定义为该文件存在副本。请返回任一存在副本的文件 id。

- 标签：数组
- 解题思路：哈希表
*/
func findRepeatDocument(documents []int) int {
	hash := make(map[int]bool)
	for _, item := range documents {
		_, ok := hash[item]
		if ok {
			return item
		} else {
			hash[item] = true
		}
	}
	return -1
}
