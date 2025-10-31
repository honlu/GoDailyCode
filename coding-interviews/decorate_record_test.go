package main

import (
	"fmt"
	"reflect"
	"testing"
)

// 测试 levelOrder 函数（普通层序遍历）
func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name     string
		input    []int // 层序遍历的数组，-1表示空节点
		expected [][]int
	}{
		{
			name:     "空树",
			input:    []int{},
			expected: [][]int(nil),
		},
		{
			name:     "单节点树",
			input:    []int{1},
			expected: [][]int{{1}},
		},
		{
			name:     "完全二叉树",
			input:    []int{3, 9, 20, -1, -1, 15, 7},
			expected: [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:     "左斜树",
			input:    []int{1, 2, -1, 3, -1, -1, -1},
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "右斜树",
			input:    []int{1, -1, 2, -1, -1, -1, 3},
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "复杂树",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			result := levelOrder(root)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("levelOrder() = %v, expected %v", result, tt.expected)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

// 基准测试
func BenchmarkLevelOrder(b *testing.B) {
	// 构建一个较大的树进行性能测试
	vals := make([]int, 100)
	for i := 0; i < 100; i++ {
		vals[i] = i + 1
	}
	root := buildTree(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		levelOrder(root)
	}
}

// 示例测试
func ExamplelevelOrder() {
	// 构建树: [3,9,20,null,null,15,7]
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	result := levelOrder(root)
	fmt.Println(result)
	// Output: [[3] [9 20] [15 7]]
}
