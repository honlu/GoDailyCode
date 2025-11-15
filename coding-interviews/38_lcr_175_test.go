package main

import (
	"testing"
)

func TestCalculateDepth(t *testing.T) {
	tests := []struct {
		name     string
		treeVals []int // 层序遍历数组，-1表示空节点
		expected int   // 期望的深度
	}{
		{
			name:     "空树",
			treeVals: []int{},
			expected: 0,
		},
		{
			name:     "单节点",
			treeVals: []int{1},
			expected: 1,
		},
		{
			name:     "基本功能 - 简单树",
			treeVals: []int{4, 2, 5, 1, 3, -1, -1},
			expected: 3,
		},
		{
			name:     "只有左子树",
			treeVals: []int{3, 2, -1, 1, -1, -1, -1},
			expected: 3,
		},
		{
			name:     "只有右子树",
			treeVals: []int{1, -1, 2, -1, -1, -1, 3},
			expected: 3,
		},
		{
			name:     "复杂树 - 完整二叉树",
			treeVals: []int{4, 2, 6, 1, 3, 5, 7},
			expected: 3,
		},
		{
			name:     "复杂树 - 不平衡树",
			treeVals: []int{5, 3, 8, 2, 4, -1, 9, 1, -1, -1, -1, -1, -1, -1, 10},
			expected: 4,
		},
		{
			name:     "深度为2的树",
			treeVals: []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "大值树",
			treeVals: []int{10, 5, 15, 3, 7, 12, 18},
			expected: 3,
		},
		{
			name:     "所有节点值相同",
			treeVals: []int{2, 2, 2},
			expected: 2,
		},
		{
			name:     "深度为4的树",
			treeVals: []int{1, 2, 3, 4, -1, -1, -1, 5, -1, -1, -1, -1, -1, -1, -1},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 构建树
			root := buildTree(tt.treeVals)

			// 调用函数
			result := calculateDepth(root)

			// 验证结果
			if result != tt.expected {
				t.Errorf("calculateDepth() = %v, expected %v", result, tt.expected)
				t.Errorf("Tree: %v", tt.treeVals)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

