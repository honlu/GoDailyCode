package misc

/*
97-169.多数投票
https://leetcode.cn/problems/majority-element/submissions/?envType=study-plan-v2&envId=top-100-liked
*/
/*
思路：
1. 排序+中位数
2. 哈希奇数
3. Boyer–Moore Majority Vote Algorithm， 多数投票奇数法。时间复杂度O（n） ✅
*/
func majorityElement(nums []int) int {
	var res, count int
	for _, item := range nums {
		if count == 0 { // item是再次遇到新值，个数为1
			res, count = item, 1
		} else if item == res { // 相同值，个数加一
			count++
		} else { // 不同值，个数减一
			count--
		}
	}
	return res
}
