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
	暴力解法：O(n2),依次遍历每一个加油站作为起点，模拟一圈。【会超时】

	方法一：
	直接从全局进行贪心选择，情况如下：
	情况一：如果gas的总和小于cost总和，那么无论从哪里出发，一定是跑不了一圈的
	情况二：rest[i] = gas[i]-cost[i]为一天剩下的油，i从0开始计算累加到最后一站，
		如果累加没有出现负数，说明从0出发，油就没有断过，那么0就是起点。
	情况三：如果累加的最小值是负数，汽车就要从非0节点出发，从后向前，
		看哪个节点能把这个负数填平，能把这个负数填平的节点就是出发节点。
	【情况3不是很好理解，要细想。】

*/

/* 超时！
暴力解法：如果跑了一圈，中途没有断油，而且最后油量大于等于0，说明这个起点是ok的。
思路比较简单，关键要写模拟跑一圈的过程代码。
*/
// func canCompleteCircuit(gas []int, cost []int) int {
// 	for i := 0; i < len(gas); i++ {
// 		res := gas[i] - cost[i] // 记录剩余油量
// 		index := (i + 1) % len(gas)
// 		for res > 0 && index != i { // 模拟以i为起点行驶一圈。如果有rest==0，那么答案就不唯一了
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

/*
方法一:
*/
func canCompleteCircuit(gas []int, cost []int) int {
	curSum := 0
	min := math.MaxInt
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		if curSum < min {
			min = curSum
		}
	}
	if curSum < 0 { // 情况1
		return -1
	}
	if min >= 0 { // 情况2
		return 0
	}
	// 情况3：min < 0时，要换成其他的节点
	for i := len(gas) - 1; i > 0; i-- {
		min += gas[i] - cost[i]
		if min >= 0 {
			return i
		}
	}
	return -1

}

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
	index := 0    // 起点
	for i := 0; i < len(gas); i++ {
		curSum += gas[i] - cost[i]
		totalSum += gas[i] - cost[i]
		if curSum < 0 { // 如果当前累加剩余油量和curSum小于0，则要更新开始起点以及剩余累加值重新为0
			index = i + 1
			curSum = 0
		}
	}
	if totalSum < 0 {
		return -1
	}
	return index
}