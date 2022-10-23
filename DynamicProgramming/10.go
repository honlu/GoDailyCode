/*
10、分割等和子集
2022-10-22
link: https://leetcode.cn/problems/partition-equal-subset-sum/
question:
	给你一个只包含正整数的非空数组nums。请你判断是否可以将这个数组分割成两个子集，
	使得两个子集的元素和相等。
answer:
	只要找到集合里能够出现sum/2的子集总和，就算是可以分割成两个相同元素和子集了。
	解法1：回溯暴力搜索，但可能超时，需要优化。
	解法2：动态规划-01背包。
	注意四点，将01背包套入本题：
		背包的体积为sum / 2.
		背包要放入的商品（集合里的元素）重量为 元素的数值，价值也为元素的数值
		背包如果正好装满，说明找到了总和为 sum / 2 的子集。
		背包中每一个元素是不可重复放入。
	动态五部曲：
	1、确定dp数组以及含义：dp[j]表示背包总容量是j，最大可以凑成j的子集总和为dp[j]。
	2、递推公式：dp[j] = max(dp[j], dp[j - weight[i]] + value[i]);
	3、dp数组初始化：初始化为0
	4、遍历顺序：滚动数组，如果使用一维dp数组，物品遍历的for循环放在外层，遍历背包的for循环放在内层，且内层for循环倒序遍历！
	5、推导dp数组，验证。

*/
// 时间复杂度：O(n^2)	空间复杂度：O(n)
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// 如果 nums的总和为奇数则不可能平分成两个子集
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	// 定义dp数组和初始化
	dp := make([]int, target+1)
	for _, v := range nums {
		for j := target; j >= v; j-- {
			if dp[j-v]+v > dp[j] {
				dp[j] = dp[j-v] + v
			}
		}
	}
	return dp[target] == target
}