package hash349

func intersection(nums1 []int, nums2 []int) []int {
	seen := map[int]bool{}
	for _, num := range nums1 {
		seen[num] = true
	}

	res := []int{}
	for _, num := range nums2 {
		if seen[num] {
			res = append(res, num)
			delete(seen, num)
		}
	}
	return res
}
