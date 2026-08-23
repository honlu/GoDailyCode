package lcr170

import "testing"

func TestReversePairs(t *testing.T) {
	tests := []struct {
		name     string
		record   []int
		expected int
	}{
		{
			name:     "官方示例",
			record:   []int{9, 7, 5, 4, 6},
			expected: 8, // (9,7)(9,5)(9,4)(9,6)(7,5)(7,4)(7,6)(5,4)
		},
		{
			name:     "剑指Offer常见示例",
			record:   []int{7, 5, 6, 4},
			expected: 5, // (7,5)(7,6)(7,4)(5,4)(6,4)
		},
		{
			name:     "空数组",
			record:   []int{},
			expected: 0,
		},
		{
			name:     "单元素",
			record:   []int{1},
			expected: 0,
		},
		{
			name:     "两个元素升序",
			record:   []int{1, 2},
			expected: 0,
		},
		{
			name:     "两个元素降序",
			record:   []int{2, 1},
			expected: 1,
		},
		{
			name:     "已排序无逆序对",
			record:   []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "完全逆序",
			record:   []int{5, 4, 3, 2, 1},
			expected: 10, // C(5,2) = 10
		},
		{
			name:     "全部相等",
			record:   []int{3, 3, 3, 3},
			expected: 0,
		},
		{
			name:     "部分相等夹杂逆序",
			record:   []int{3, 1, 3, 1},
			expected: 3, // (3,1)(3,1)(3,1)
		},
		{
			name:     "含负数",
			record:   []int{1, -1, 0, -2},
			expected: 5, // (1,-1)(1,0)(1,-2)(-1,-2)(0,-2)
		},
		{
			name:     "仅右半段有逆序",
			record:   []int{1, 2, 4, 3},
			expected: 1,
		},
		{
			name:     "跨越左右半段的逆序",
			record:   []int{3, 1, 2},
			expected: 2, // (3,1)(3,2)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			record := append([]int(nil), tt.record...)
			got := reversePairs(record)
			if got != tt.expected {
				t.Errorf("reversePairs(%v) = %d, expected %d", tt.record, got, tt.expected)
			}
		})
	}
}

func TestReversePairsMatchesBruteForce(t *testing.T) {
	cases := [][]int{
		{1, 3, 2, 3, 1},
		{0, 0, 0},
		{-5, -1, -3, -2},
		{8, 4, 2, 1, 16, 9},
		{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	}

	for _, record := range cases {
		t.Run("", func(t *testing.T) {
			want := countReversePairsBrute(record)
			got := reversePairs(append([]int(nil), record...))
			if got != want {
				t.Errorf("reversePairs(%v) = %d, brute force = %d", record, got, want)
			}
		})
	}
}

func countReversePairsBrute(record []int) int {
	cnt := 0
	for i := 0; i < len(record); i++ {
		for j := i + 1; j < len(record); j++ {
			if record[i] > record[j] {
				cnt++
			}
		}
	}
	return cnt
}
