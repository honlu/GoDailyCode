package monotonicstack3

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}

	stack := []int{}
	for i := 0; i < len(nums)*2; i++ {
		idx := i % len(nums)
		for len(stack) > 0 && nums[idx] > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] = nums[idx]
		}
		if i < len(nums) {
			stack = append(stack, idx)
		}
	}
	return res
}
