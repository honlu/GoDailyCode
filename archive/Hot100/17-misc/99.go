package misc

/*
99-31.下一个排列
https://leetcode.cn/problems/next-permutation/?envType=study-plan-v2&envId=top-100-liked
*/
/*
思路：数组转化的规律。（比较难梳理通）
*/
func nextPermutation(nums []int) {
	// 首先从末尾找逆序排列的范围
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] { // (i+1,n)是降序排列，选择i作为交换值
		i--
	}
	if i >= 0 { // 在逆序排列中，倒着找到比nums[i]第一个大的nums[j]，然后交换。则[i+1,len(num)-1)仍然是降序
		for j := len(nums) - 1; j >= i; j-- {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
	// 将[i+1,len(nums)-1]翻转改为升序
	for i, j := i+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return
}
