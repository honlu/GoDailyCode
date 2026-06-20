package hot100

/*
17-缺失的第一个正数.
*/
func firstMissingPositive(nums []int) int {
	max := 0
	numMap := make(map[int]bool)
	for _, num := range nums {
		if num > max {
			max = num
		}
		numMap[num] = true
	}
	i := 1
	for ; i <= max; i++ {
		if !numMap[i] {
			return i
		}
	}
	return i
}
