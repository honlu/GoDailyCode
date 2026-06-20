package main

// map实现
func trainingPlan(actions []int) int {
	tmpMap := make(map[int]int)
	for _, item := range actions {
		tmpMap[item]++
	}
	for k, v := range tmpMap {
		if v == 1 {
			return k
		}
	}
	return -1
}

// 待使用位运算改进
