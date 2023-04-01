/*
剑指offer2-61. 扑克牌中的顺子
2023-4-1 by lu
link: https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
question:
	从若干副扑克牌中随机抽 5 张牌，判断是不是一个顺子，
	即这5张牌是不是连续的。2～10为数字本身，
	A为1，J为11，Q为12，K为13，而大、小王为 0 ，
	可以看成任意数字。A 不能视为 14。
answer:
	若干副扑克牌，所以王的个数不只两个。小心这个细节
*/
func isStraight(nums []int) bool {
	// 若干副牌，所以大小王不仅仅两个
	sort.Ints(nums)
	zero := 0
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			zero++
		} else if nums[i] == nums[i+1] {
			return false
		}
	}
	return nums[4]-nums[zero] < 5 // 如果3个王，就不能保证最大值和最小值等于4了，所以和5大小判断
}