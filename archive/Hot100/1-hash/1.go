package hot100

/*
1. 两数之和
*/

func twoSum(nums []int, target int) []int {
	// 临时map k-v为nums:值-索引，用于快速查找目标值的补数是否已存在
	valueIndexMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := valueIndexMap[target-num]; ok {
			return []int{i, j}
		} else {
			valueIndexMap[num] = i
		}
	}
	return []int{}
}
