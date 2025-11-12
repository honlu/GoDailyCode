package main

import (
	"testing"
)

func TestTreeToDoublyList(t *testing.T) {
	tests := []struct {
		name     string
		treeVals []int // 层序遍历数组，-1表示空节点
		expected []int // 期望的双向链表前向遍历结果
	}{
		{
			name:     "空树",
			treeVals: []int{},
			expected: nil,
		},
		{
			name:     "单节点",
			treeVals: []int{1},
			expected: []int{1},
		},
		{
			name:     "基本功能 - 简单树",
			treeVals: []int{4, 2, 5, 1, 3, -1, -1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "只有左子树",
			treeVals: []int{3, 2, -1, 1, -1, -1, -1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "只有右子树",
			treeVals: []int{1, -1, 2, -1, -1, -1, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "复杂树 - 完整二叉树",
			treeVals: []int{4, 2, 6, 1, 3, 5, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "复杂树 - 不平衡树",
			treeVals: []int{5, 3, 8, 2, 4, -1, 9, 1, -1, -1, -1, -1, -1, -1, 10},
			expected: []int{1, 2, 3, 4, 5, 8, 9, 10},
		},
		{
			name:     "所有节点值相同",
			treeVals: []int{2, 2, 2},
			expected: []int{2, 2, 2},
		},
		{
			name:     "大值树",
			treeVals: []int{10, 5, 15, 3, 7, 12, 18},
			expected: []int{3, 5, 7, 10, 12, 15, 18},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 构建树
			root := buildTree(tt.treeVals)

			// 转换为双向链表
			head := treeToDoublyList(root)

			// 验证结果
			if tt.expected == nil {
				if head != nil {
					t.Errorf("treeToDoublyList() = %v, expected nil", head)
				}
				return
			}

			if head == nil {
				t.Errorf("treeToDoublyList() = nil, expected non-nil")
				return
			}

			// 前向遍历验证（Right 指针）
			forwardVals := make([]int, 0)
			current := head
			visited := make(map[*TreeNode]bool)
			for current != nil && !visited[current] {
				forwardVals = append(forwardVals, current.Val)
				visited[current] = true
				current = current.Right
			}

			if len(forwardVals) != len(tt.expected) {
				t.Errorf("前向遍历长度不匹配: got %v, expected %v", forwardVals, tt.expected)
				return
			}

			for i := 0; i < len(forwardVals); i++ {
				if forwardVals[i] != tt.expected[i] {
					t.Errorf("前向遍历值不匹配: got %v, expected %v", forwardVals, tt.expected)
					return
				}
			}

			// 后向遍历验证（Left 指针）
			// 找到尾节点
			tail := head
			for tail.Right != nil && tail.Right != head {
				tail = tail.Right
			}

			backwardVals := make([]int, 0)
			current = tail
			visited = make(map[*TreeNode]bool)
			for current != nil && !visited[current] {
				backwardVals = append(backwardVals, current.Val)
				visited[current] = true
				current = current.Left
			}

			// 反转期望值用于后向验证
			expectedBackward := make([]int, len(tt.expected))
			for i := 0; i < len(tt.expected); i++ {
				expectedBackward[i] = tt.expected[len(tt.expected)-1-i]
			}

			if len(backwardVals) != len(expectedBackward) {
				t.Errorf("后向遍历长度不匹配: got %v, expected %v", backwardVals, expectedBackward)
				return
			}

			for i := 0; i < len(backwardVals); i++ {
				if backwardVals[i] != expectedBackward[i] {
					t.Errorf("后向遍历值不匹配: got %v, expected %v", backwardVals, expectedBackward)
					return
				}
			}

			t.Logf("✓ Test passed: %s", tt.name)
		})
	}
}
