package p384

import "math/rand"

type Solution struct {
	nums   []int
	origin []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums:   append([]int{}, nums...),
		origin: append([]int{}, nums...),
	}
}

func (s *Solution) Reset() []int {
	s.nums = append([]int{}, s.origin...)
	return append([]int{}, s.nums...)
}

func (s *Solution) Shuffle() []int {
	res := append([]int{}, s.nums...)
	for i := len(res) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		res[i], res[j] = res[j], res[i]
	}
	s.nums = append([]int{}, res...)
	return res
}
