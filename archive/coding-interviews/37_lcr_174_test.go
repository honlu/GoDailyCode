package main

import (
	"testing"
)

func TestFindTargetNode(t *testing.T) {
	tests := []struct {
		name     string
		treeVals []int // 层序遍历数组，-1表示空节点
		cnt      int   // 倒数第几个
		expected int   // 期望的结果
	}{
		{
			name:     "单节点 - cnt=1",
			treeVals: []int{1},
			cnt:      1,
			expected: 1,
		},
		{
			name:     "基本功能 - 简单树，cnt=1（最后一个）",
			treeVals: []int{4, 2, 5, 1, 3, -1, -1},
			cnt:      1,
			expected: 5,
		},
		{
			name:     "基本功能 - 简单树，cnt=2",
			treeVals: []int{4, 2, 5, 1, 3, -1, -1},
			cnt:      2,
			expected: 4,
		},
		{
			name:     "基本功能 - 简单树，cnt=3",
			treeVals: []int{4, 2, 5, 1, 3, -1, -1},
			cnt:      3,
			expected: 3,
		},
		{
			name:     "基本功能 - 简单树，cnt=4",
			treeVals: []int{4, 2, 5, 1, 3, -1, -1},
			cnt:      4,
			expected: 2,
		},
		{
			name:     "基本功能 - 简单树，cnt=5（第一个）",
			treeVals: []int{4, 2, 5, 1, 3, -1, -1},
			cnt:      5,
			expected: 1,
		},
		{
			name:     "只有左子树 - cnt=1",
			treeVals: []int{3, 2, -1, 1, -1, -1, -1},
			cnt:      1,
			expected: 3,
		},
		{
			name:     "只有左子树 - cnt=2",
			treeVals: []int{3, 2, -1, 1, -1, -1, -1},
			cnt:      2,
			expected: 2,
		},
		{
			name:     "只有左子树 - cnt=3",
			treeVals: []int{3, 2, -1, 1, -1, -1, -1},
			cnt:      3,
			expected: 1,
		},
		{
			name:     "只有右子树 - cnt=1",
			treeVals: []int{1, -1, 2, -1, -1, -1, 3},
			cnt:      1,
			expected: 3,
		},
		{
			name:     "只有右子树 - cnt=2",
			treeVals: []int{1, -1, 2, -1, -1, -1, 3},
			cnt:      2,
			expected: 2,
		},
		{
			name:     "只有右子树 - cnt=3",
			treeVals: []int{1, -1, 2, -1, -1, -1, 3},
			cnt:      3,
			expected: 1,
		},
		{
			name:     "复杂树 - 完整二叉树，cnt=1",
			treeVals: []int{4, 2, 6, 1, 3, 5, 7},
			cnt:      1,
			expected: 7,
		},
		{
			name:     "复杂树 - 完整二叉树，cnt=4",
			treeVals: []int{4, 2, 6, 1, 3, 5, 7},
			cnt:      4,
			expected: 4,
		},
		{
			name:     "复杂树 - 完整二叉树，cnt=7",
			treeVals: []int{4, 2, 6, 1, 3, 5, 7},
			cnt:      7,
			expected: 1,
		},
		{
			name:     "复杂树 - 不平衡树，cnt=1",
			treeVals: []int{5, 3, 8, 2, 4, -1, 9, 1, -1, -1, -1, -1, -1, -1, 10},
			cnt:      1,
			expected: 10,
		},
		{
			name:     "复杂树 - 不平衡树，cnt=5",
			treeVals: []int{5, 3, 8, 2, 4, -1, 9, 1, -1, -1, -1, -1, -1, -1, 10},
			cnt:      5,
			expected: 4,
		},
		{
			name:     "复杂树 - 不平衡树，cnt=8",
			treeVals: []int{5, 3, 8, 2, 4, -1, 9, 1, -1, -1, -1, -1, -1, -1, 10},
			cnt:      8,
			expected: 1,
		},
		{
			name:     "大值树 - cnt=1",
			treeVals: []int{10, 5, 15, 3, 7, 12, 18},
			cnt:      1,
			expected: 18,
		},
		{
			name:     "大值树 - cnt=4",
			treeVals: []int{10, 5, 15, 3, 7, 12, 18},
			cnt:      4,
			expected: 10,
		},
		{
			name:     "大值树 - cnt=7",
			treeVals: []int{10, 5, 15, 3, 7, 12, 18},
			cnt:      7,
			expected: 3,
		},
		{
			name:     "所有节点值相同 - cnt=1",
			treeVals: []int{2, 2, 2},
			cnt:      1,
			expected: 2,
		},
		{
			name:     "所有节点值相同 - cnt=2",
			treeVals: []int{2, 2, 2},
			cnt:      2,
			expected: 2,
		},
		{
			name:     "所有节点值相同 - cnt=3",
			treeVals: []int{2, 2, 2},
			cnt:      3,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 构建树
			root := buildTree(tt.treeVals)

			// 调用函数
			result := findTargetNode(root, tt.cnt)

			// 验证结果
			if result != tt.expected {
				t.Errorf("findTargetNode() = %v, expected %v", result, tt.expected)
				t.Errorf("Tree: %v, cnt: %d", tt.treeVals, tt.cnt)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

