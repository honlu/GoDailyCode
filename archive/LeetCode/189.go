/*
189. 轮转数组
2023-3-10
link: https://leetcode.cn/problems/rotate-array/
question:
	给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
answer:
	不同的方式，要根据不同语言的特性即语法的底层结构进行实验。
*/
// 第二种 利用O(k)的额外数组
func rotate(nums []int, k int) {
	// 第一种,仅仅append改变不了外部切片指向
	// nums = append(nums[k:], append([]int{}, nums[:k]...)...)

	if k > len(nums) {
		k = k % len(nums)
	}
	// 第二种
	temp := append([]int{}, nums[len(nums)-k:]...) // 注意不可以直接temp := nums[len(nums)-k:],否则nums指向的值更改，temp也会更改
	for i := len(nums) - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}
	for i := 0; i < k; i++ {
		nums[i] = temp[i]
	}
}

// 第三种
func rotate(nums []int, k int) {
	// 第三种：反转数组
	if k > len(nums) {
		k = k % len(nums)
	}
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}