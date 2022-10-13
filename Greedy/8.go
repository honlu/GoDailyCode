/*
8、K次取反后最大化的数组和
2022-10-12
link:1005-https://leetcode.cn/problems/maximize-sum-of-array-after-k-negations/
question:
	给定一个整数数组 A，和要修改的次数K。我们只能用以下方法修改该数组：
	我们选择某个索引 i 并将 A[i] 替换为 -A[i]，然后总共重复这个过程 K 次。
	（我们可以多次选择同一个索引 i。）
	以这种方式修改数组后，返回数组可能的最大和。
answer:
	贪心思想
	局部最优：让绝对值大的负数变为正数，当前数值达到最大，整体最优：整个数组和达到最大。
	局部最优可以推出全局最优。

	如何转变K次正负，让 数组和 达到最大。
	又是一个贪心：局部最优：只找数值最小的正整数进行反转，当前数值可以达到最大
	（例如正整数数组{5, 3, 1}，反转1 得到-1 比 反转5得到的-5 大多了），
	全局最优：整个 数组和 达到最大。

	接替步骤：
	第一步：将数组按照绝对值大小从大到小排序，注意要按照绝对值的大小
	第二步：从前向后遍历，遇到负数将其变为正数，同时K--
	第三步：如果K还大于0，那么反复转变数值最小的元素，将K用完
	第四步：求和
*/
package main

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	// 1、排序_最重要
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	}) // 逆序排序。正数按照从大到小排序，负数按照从小到大排序
	// 2、从前往后遍历
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	if k%2 == 1 { // 3、如果k还大于0,则反转最小的元素，k奇数取反，k偶数反转后不变
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	// 4、求和
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}
