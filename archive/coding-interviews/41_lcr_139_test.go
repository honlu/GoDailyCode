package main

import (
	"testing"
)

// findNode 在BST中查找值为val的节点
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if root.Val > val {
		return findNode(root.Left, val)
	}
	return findNode(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name     string
		treeVals []int // 层序遍历数组，-1表示空节点（必须是有效的BST）
		pVal     int   // 节点p的值
		qVal     int   // 节点q的值
		expected int   // 期望的LCA节点值
	}{
		{
			name:     "基本功能 - p和q在不同子树",
			treeVals: []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5},
			pVal:     2,
			qVal:     8,
			expected: 6,
		},
		{
			name:     "基本功能 - p和q在左子树",
			treeVals: []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5},
			pVal:     0,
			qVal:     4,
			expected: 2,
		},
		{
			name:     "基本功能 - p和q在右子树",
			treeVals: []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5},
			pVal:     7,
			qVal:     9,
			expected: 8,
		},
		{
			name:     "p是q的祖先",
			treeVals: []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5},
			pVal:     2,
			qVal:     4,
			expected: 2,
		},
		{
			name:     "q是p的祖先",
			treeVals: []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5},
			pVal:     3,
			qVal:     2,
			expected: 2,
		},
		{
			name:     "根节点是LCA",
			treeVals: []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5},
			pVal:     0,
			qVal:     9,
			expected: 6,
		},
		{
			name:     "p和q是同一个节点",
			treeVals: []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5},
			pVal:     4,
			qVal:     4,
			expected: 4,
		},
		{
			name:     "简单BST - 三个节点",
			treeVals: []int{2, 1, 3},
			pVal:     1,
			qVal:     3,
			expected: 2,
		},
		{
			name:     "简单BST - p是根",
			treeVals: []int{2, 1, 3},
			pVal:     2,
			qVal:     3,
			expected: 2,
		},
		{
			name:     "简单BST - q是根",
			treeVals: []int{2, 1, 3},
			pVal:     1,
			qVal:     2,
			expected: 2,
		},
		{
			name:     "单节点树",
			treeVals: []int{1},
			pVal:     1,
			qVal:     1,
			expected: 1,
		},
		{
			name:     "复杂BST - 深度为4",
			treeVals: []int{10, 5, 15, 3, 7, 12, 18, 2, 4, 6, 8},
			pVal:     2,
			qVal:     8,
			expected: 5,
		},
		{
			name:     "复杂BST - 左子树深处",
			treeVals: []int{10, 5, 15, 3, 7, 12, 18, 2, 4, 6, 8},
			pVal:     2,
			qVal:     4,
			expected: 3,
		},
		{
			name:     "复杂BST - 右子树深处",
			treeVals: []int{10, 5, 15, 3, 7, 12, 18, 2, 4, 6, 8},
			pVal:     12,
			qVal:     18,
			expected: 15,
		},
		{
			name:     "复杂BST - 跨子树",
			treeVals: []int{10, 5, 15, 3, 7, 12, 18, 2, 4, 6, 8},
			pVal:     4,
			qVal:     12,
			expected: 10,
		},
		{
			name:     "不平衡BST - 只有右子树",
			treeVals: []int{1, -1, 2, -1, -1, -1, 3},
			pVal:     2,
			qVal:     3,
			expected: 2,
		},
		{
			name:     "不平衡BST - 只有左子树",
			treeVals: []int{3, 2, -1, 1, -1, -1, -1},
			pVal:     1,
			qVal:     2,
			expected: 2,
		},
		{
			name:     "大值BST",
			treeVals: []int{100, 50, 150, 25, 75, 125, 175},
			pVal:     25,
			qVal:     175,
			expected: 100,
		},
		{
			name:     "相邻节点",
			treeVals: []int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5},
			pVal:     3,
			qVal:     5,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 构建BST
			root := buildTree(tt.treeVals)

			// 查找节点p和q
			p := findNode(root, tt.pVal)
			q := findNode(root, tt.qVal)

			// 验证节点存在
			if p == nil {
				t.Fatalf("节点p (值=%d) 在树中不存在", tt.pVal)
			}
			if q == nil {
				t.Fatalf("节点q (值=%d) 在树中不存在", tt.qVal)
			}

			// 调用函数
			result := lowestCommonAncestor(root, p, q)

			// 验证结果
			if result == nil {
				t.Errorf("lowestCommonAncestor() = nil, expected node with value %d", tt.expected)
				t.Errorf("Tree: %v, p: %d, q: %d", tt.treeVals, tt.pVal, tt.qVal)
			} else if result.Val != tt.expected {
				t.Errorf("lowestCommonAncestor() = %d, expected %d", result.Val, tt.expected)
				t.Errorf("Tree: %v, p: %d, q: %d", tt.treeVals, tt.pVal, tt.qVal)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

