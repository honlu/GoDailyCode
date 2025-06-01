package hot100

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			nums:     []int{-2, 0, 1, 1, 2},
			expected: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			nums:     []int{},
			expected: [][]int{},
		},
		{
			nums:     []int{0},
			expected: [][]int{},
		},
	}

	for _, test := range tests {
		result := threeSum(test.nums)
		if !compareTriplets(result, test.expected) {
			t.Errorf("threeSum(%v) = %v; expected %v", test.nums, result, test.expected)
		} else {
			t.Logf("threeSum(%v) passed", test.nums)
		}
	}
}

func compareTriplets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	visited := make([]bool, len(b))
	for _, tripletA := range a {
		sort.Ints(tripletA)
		found := false
		for j, tripletB := range b {
			sort.Ints(tripletB)
			if !visited[j] && reflect.DeepEqual(tripletA, tripletB) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
