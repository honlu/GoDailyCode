package main

import (
	"testing"
)

func TestCheckSymmetricTree(t *testing.T) {
	tests := []struct {
		name     string
		input    []int // 树的层序遍历数组，-1表示空节点
		expected bool
	}{
		{
			name:     "空树",
			input:    []int{},
			expected: true,
		},
		{
			name:     "单节点树",
			input:    []int{1},
			expected: true,
		},
		{
			name:     "对称二叉树 - 两个子节点",
			input:    []int{1, 2, 2},
			expected: true,
		},
		{
			name:     "对称二叉树 - 完全对称",
			input:    []int{1, 2, 2, 3, 4, 4, 3},
			expected: true,
		},
		{
			name:     "对称二叉树 - 多层对称",
			input:    []int{1, 2, 2, 3, -1, -1, 3},
			expected: true,
		},
		{
			name:     "不对称二叉树 - 值不相等",
			input:    []int{1, 2, 3},
			expected: false,
		},
		{
			name:     "不对称二叉树 - 结构不对称",
			input:    []int{1, 2, 2, -1, 3, -1, 3},
			expected: false,
		},
		{
			name:     "不对称二叉树 - 只有左子树",
			input:    []int{1, 2, -1, 3},
			expected: false,
		},
		{
			name:     "不对称二叉树 - 只有右子树",
			input:    []int{1, -1, 2, -1, -1, -1, 3},
			expected: false,
		},
		{
			name:     "对称二叉树 - 复杂结构（一侧为空）",
			input:    []int{1, 2, 2, -1, 3, 3, -1},
			expected: true,
		},
		{
			name:     "对称二叉树 - 更深层对称",
			input:    []int{1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 8, 8, 7, 6, 5},
			expected: true,
		},
		{
			name:     "不对称二叉树 - 深层值不相等",
			input:    []int{1, 2, 2, 3, 4, 4, 3, 5, 6, 7, 8, 8, 7, 6, 9},
			expected: false,
		},
		{
			name:     "不对称二叉树 - 单侧不对称",
			input:    []int{1, 2, 2, 3, -1, 3, -1},
			expected: false,
		},
		{
			name:     "对称二叉树 - 所有节点值相同但结构对称",
			input:    []int{1, 1, 1, 1, 1, 1, 1},
			expected: true,
		},
		{
			name:     "不对称二叉树 - 单侧不对称（左节点的左子树和右节点的左子树）",
			input:    []int{1, 2, 2, 3, 3},
			expected: false,
		},
		{
			name:     "不对称二叉树 - 值不相等但结构对称",
			input:    []int{1, 2, 2, 3, 4},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			result := checkSymmetricTree(root)

			if result != tt.expected {
				t.Errorf("checkSymmetricTree() = %v, expected %v", result, tt.expected)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}
