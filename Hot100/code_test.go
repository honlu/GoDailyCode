package hot100

import (
	"reflect"
	"sort"
	"testing"
)

// 题1的测试用例
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

// 题2的测试用例
func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected [][]string
	}{
		{
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
		{
			input:    []string{""},
			expected: [][]string{{""}},
		},
		{
			input:    []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for _, test := range tests {
		result := groupAnagrams(test.input)
		if !compareAnagramGroups(result, test.expected) {
			t.Errorf("Test failed for input %v. Got %v, expected %v", test.input, result, test.expected)
		} else {
			t.Logf("Test passed for input %v", test.input)
		}
	}
}

func compareAnagramGroups(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}
	for _, groupA := range a {
		found := false
		for _, groupB := range b {
			if reflect.DeepEqual(groupA, groupB) {
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
