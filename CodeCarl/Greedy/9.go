/*
9、加油站
2022-10-12
link:134-https://leetcode.cn/problems/gas-station/
question:
	在一条环路上有 N 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
	你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。
	你从其中的一个加油站出发，开始时油箱为空。
	如果你可以绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1。

answer:
	暴力解法：O(n2),依次遍历每一个加油站作为起点，模拟一圈。

*/

/* 超时！
暴力解法：如果跑了一圈，中途没有断油，而且最后油量大于等于0，说明这个起点是ok的。
思路比较简单，关键要写模拟跑一圈的过程代码。
*/
// func canCompleteCircuit(gas []int, cost []int) int {
// 	for i := 0; i < len(gas); i++ {
// 		res := gas[i] - cost[i] // 记录剩余油量
// 		index := (i + 1) % len(gas)
// 		for res > 0 && index != i { // 模拟以i为起点行驶一圈
// 			res += gas[index] - cost[index]
// 			index = (i + 1) % len(gas)
// 		}
// 		// 如果以i为起点跑一圈，剩余油量>=0,返回该起始位置
// 		if res >= 0 && index == i {
// 			return i
// 		}
// 	}
// 	return -1
// }

// 贪心算法——方法2
/*
思路：
首先:如果总油量减去总消耗大于等于零那么一定可以跑完一圈，
	说明各个站点的加油站 剩油量rest[i]相加一定是大于等于零的。
其次：i从0开始累加rest[i]，和记为curSum，一旦curSum小于零，
说明[0, i]区间都不能作为起始位置，起始位置从i+1算起，再从0计算curSum。
*/
func canCompleteCircuit(gas []int, cost []int) int {
	curSum := 0   // 当前节点剩余油量
	totalSum := 0 // 总共剩余油量=总油量-总消耗量
	start := 0    // 起点
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 { // ?
			start = i + 1
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return start
}