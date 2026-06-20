package main

import (
	"testing"
)

// findNodeInBT 在普通二叉树中查找值为val的节点（使用DFS）
func findNodeInBT(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	// 先查找左子树
	if left := findNodeInBT(root.Left, val); left != nil {
		return left
	}
	// 再查找右子树
	return findNodeInBT(root.Right, val)
}

func TestLowestCommonAncestor194(t *testing.T) {
	tests := []struct {
		name     string
		treeVals []int // 层序遍历数组，-1表示空节点
		pVal     int   // 节点p的值
		qVal     int   // 节点q的值
		expected int   // 期望的LCA节点值
	}{
		{
			name:     "基本功能 - p和q在不同子树",
			treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4},
			pVal:     5,
			qVal:     1,
			expected: 3,
		},
		{
			name:     "基本功能 - p和q在左子树",
			treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4},
			pVal:     6,
			qVal:     2,
			expected: 5,
		},
		{
			name:     "基本功能 - p和q在右子树",
			treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4},
			pVal:     0,
			qVal:     8,
			expected: 1,
		},
		{
			name:     "p是q的祖先",
			treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4},
			pVal:     5,
			qVal:     4,
			expected: 5,
		},
		{
			name:     "q是p的祖先",
			treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4},
			pVal:     7,
			qVal:     5,
			expected: 5,
		},
		{
			name:     "根节点是LCA",
			treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4},
			pVal:     6,
			qVal:     0,
			expected: 3,
		},
		{
			name:     "p和q是同一个节点",
			treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4},
			pVal:     2,
			qVal:     2,
			expected: 2,
		},
		{
			name:     "简单树 - 三个节点",
			treeVals: []int{1, 2, 3},
			pVal:     2,
			qVal:     3,
			expected: 1,
		},
		{
			name:     "简单树 - p是根",
			treeVals: []int{1, 2, 3},
			pVal:     1,
			qVal:     3,
			expected: 1,
		},
		{
			name:     "简单树 - q是根",
			treeVals: []int{1, 2, 3},
			pVal:     2,
			qVal:     1,
			expected: 1,
		},
		{
			name:     "单节点树",
			treeVals: []int{1},
			pVal:     1,
			qVal:     1,
			expected: 1,
		},
		{
			name:     "复杂树 - 深度为4",
			treeVals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			pVal:     8,
			qVal:     11,
			expected: 2,
		},
		{
			name:     "复杂树 - 左子树深处",
			treeVals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			pVal:     8,
			qVal:     9,
			expected: 4,
		},
		{
			name:     "复杂树 - 右子树深处",
			treeVals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			pVal:     14,
			qVal:     15,
			expected: 7,
		},
		{
			name:     "复杂树 - 跨子树",
			treeVals: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			pVal:     9,
			qVal:     12,
			expected: 1,
		},
		{
			name:     "不平衡树 - 只有右子树",
			treeVals: []int{1, -1, 2, -1, -1, -1, 3},
			pVal:     2,
			qVal:     3,
			expected: 2,
		},
		{
			name:     "不平衡树 - 只有左子树",
			treeVals: []int{3, 2, -1, 1, -1, -1, -1},
			pVal:     1,
			qVal:     2,
			expected: 2,
		},
		{
			name:     "相邻节点",
			treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4},
			pVal:     7,
			qVal:     4,
			expected: 2,
		},
		{
			name:     "p和q在更深层",
			treeVals: []int{1, 2, 3, -1, 4, 5, 6, -1, -1, 7, 8},
			pVal:     7,
			qVal:     8,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 构建二叉树
			root := buildTree(tt.treeVals)

			// 查找节点p和q
			p := findNodeInBT(root, tt.pVal)
			q := findNodeInBT(root, tt.qVal)

			// 验证节点存在
			if p == nil {
				t.Fatalf("节点p (值=%d) 在树中不存在", tt.pVal)
			}
			if q == nil {
				t.Fatalf("节点q (值=%d) 在树中不存在", tt.qVal)
			}

			// 调用函数
			result := lowestCommonAncestor194(root, p, q)

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
