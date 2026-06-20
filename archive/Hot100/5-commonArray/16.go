package hot100

/*
16-除自身以外数组的乘积
*/
func productExceptSelf(nums []int) []int {
	// 前缀积*后缀积
	prePro := make([]int, len(nums))
	sufPro := make([]int, len(nums))
	res := make([]int, len(nums))
	prePro[0] = 1
	sufPro[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		prePro[i] = prePro[i-1] * nums[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		sufPro[i] = sufPro[i+1] * nums[i+1]
	}
	for i := 0; i < len(nums); i++ {
		res[i] = prePro[i] * sufPro[i]
	}
	return res

}
