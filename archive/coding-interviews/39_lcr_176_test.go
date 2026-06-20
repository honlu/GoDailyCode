package main

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name     string
		treeVals []int // 层序遍历数组，-1表示空节点
		expected bool  // 期望的结果
	}{
		{
			name:     "空树",
			treeVals: []int{},
			expected: true,
		},
		{
			name:     "单节点",
			treeVals: []int{1},
			expected: true,
		},
		{
			name:     "平衡树 - 简单平衡树",
			treeVals: []int{3, 9, 20, -1, -1, 15, 7},
			expected: true,
		},
		{
			name:     "平衡树 - 完整二叉树",
			treeVals: []int{4, 2, 6, 1, 3, 5, 7},
			expected: true,
		},
		{
			name:     "不平衡树 - 左右子树高度差超过1",
			treeVals: []int{1, 2, 2, 3, 3, -1, -1, 4, 4},
			expected: false, // 这个树实际上是不平衡的，因为左子树的左子树深度为3，右子树深度为1
		},
		{
			name:     "不平衡树 - 左子树过深",
			treeVals: []int{1, 2, 2, 3, -1, -1, -1, 4, -1, -1, -1, -1, -1, -1, -1},
			expected: false,
		},
		{
			name:     "不平衡树 - 右子树过深",
			treeVals: []int{1, 2, 2, -1, 3, -1, -1, -1, -1, -1, 4, -1, -1, -1, -1},
			expected: false,
		},
		{
			name:     "平衡树 - 只有左子树但深度为1",
			treeVals: []int{1, 2, -1},
			expected: true,
		},
		{
			name:     "平衡树 - 只有右子树但深度为1",
			treeVals: []int{1, -1, 2},
			expected: true,
		},
		{
			name:     "不平衡树 - 只有左子树深度为2",
			treeVals: []int{1, 2, -1, 3, -1, -1, -1},
			expected: false,
		},
		{
			name:     "不平衡树 - 只有右子树深度为2",
			treeVals: []int{1, -1, 2, -1, -1, -1, 3},
			expected: false,
		},
		{
			name:     "平衡树 - 大值树",
			treeVals: []int{10, 5, 15, 3, 7, 12, 18},
			expected: true,
		},
		{
			name:     "不平衡树 - 复杂不平衡树",
			treeVals: []int{1, 2, 2, 3, -1, -1, -1, 4, -1, -1, -1, -1, -1, -1, -1, 5, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1},
			expected: false,
		},
		{
			name:     "平衡树 - 所有节点值相同且平衡",
			treeVals: []int{2, 2, 2},
			expected: true,
		},
		{
			name:     "平衡树 - 左右子树高度差为0",
			treeVals: []int{1, 2, 3, 4, 5, 6, 7},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 构建树
			root := buildTree(tt.treeVals)

			// 调用函数
			result := isBalanced(root)

			// 验证结果
			if result != tt.expected {
				t.Errorf("isBalanced() = %v, expected %v", result, tt.expected)
				t.Errorf("Tree: %v", tt.treeVals)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

