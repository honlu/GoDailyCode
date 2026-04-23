package main

// 撞色搭配
// func sockCollocation(sockets []int) []int {
// 	// 使用map存储
// 	sockMap := make(map[int]bool)
// 	for _, item := range sockets {
// 		if sockMap[item] {
// 			delete(sockMap, item)
// 		} else {
// 			sockMap[item] = true
// 		}
// 	}
// 	var res []int
// 	for item, _ := range sockMap {
// 		a := item
// 		res = append(res, a)
// 	}
// 	return res
// }
// 采用位运算来解决题目，待改进
