package hot100

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, expected: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, expected: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, expected: []int{0, 1}},
		{nums: []int{1, 2, 3, 4}, target: 8, expected: []int{}},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		sort.Ints(result) // 结果顺序化
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("twoSum(%v, %d) = %v; expected %v", test.nums, test.target, result, test.expected)
		} else {
			t.Logf("twoSum(%v, %d) passed", test.nums, test.target)
		}
	}
}
