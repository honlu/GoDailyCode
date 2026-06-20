package main

import (
	"reflect"
	"testing"
)

// getPreorder 获取树的前序遍历
func getPreorder(root *TreeNode) []int {
	var result []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return result
}

// getInorder 获取树的中序遍历
func getInorder(root *TreeNode) []int {
	var result []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		result = append(result, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return result
}

func TestDeduceTree(t *testing.T) {
	tests := []struct {
		name           string
		preorder       []int
		inorder        []int
		expectedTree   []int // 层序遍历数组，-1表示空节点
		description    string
	}{
		{
			name:         "空树",
			preorder:     []int{},
			inorder:      []int{},
			expectedTree: []int{},
			description:  "空数组应该返回 nil",
		},
		{
			name:         "单节点",
			preorder:     []int{1},
			inorder:      []int{1},
			expectedTree: []int{1},
			description:  "单节点树",
		},
		{
			name:         "简单二叉树 - 只有左子树",
			preorder:     []int{1, 2},
			inorder:      []int{2, 1},
			expectedTree: []int{1, 2, -1},
			description:  "根节点为1，左子节点为2",
		},
		{
			name:         "简单二叉树 - 只有右子树",
			preorder:     []int{1, 3},
			inorder:      []int{1, 3},
			expectedTree: []int{1, -1, 3},
			description:  "根节点为1，右子节点为3",
		},
		{
			name:         "简单二叉树 - 左右子树",
			preorder:     []int{1, 2, 3},
			inorder:      []int{2, 1, 3},
			expectedTree: []int{1, 2, 3},
			description:  "根节点为1，左子节点为2，右子节点为3",
		},
		{
			name:         "完整二叉树",
			preorder:     []int{1, 2, 4, 5, 3, 6, 7},
			inorder:      []int{4, 2, 5, 1, 6, 3, 7},
			expectedTree: []int{1, 2, 3, 4, 5, 6, 7},
			description:  "完整二叉树，7个节点",
		},
		{
			name:         "不平衡树 - 左斜",
			preorder:     []int{1, 2, 3},
			inorder:      []int{3, 2, 1},
			expectedTree: []int{1, 2, -1, 3},
			description:  "左斜树",
		},
		{
			name:         "不平衡树 - 右斜",
			preorder:     []int{1, 2, 3},
			inorder:      []int{1, 2, 3},
			expectedTree: []int{1, -1, 2, -1, -1, -1, 3},
			description:  "右斜树",
		},
		{
			name:         "复杂树 - 多层级",
			preorder:     []int{3, 9, 20, 15, 7},
			inorder:      []int{9, 3, 15, 20, 7},
			expectedTree: []int{3, 9, 20, -1, -1, 15, 7},
			description:  "LeetCode 经典示例",
		},
		{
			name:         "复杂树 - 包含负数",
			preorder:     []int{-1, 2, -2, 4, 3},
			inorder:      []int{-2, 2, 4, -1, 3},
			expectedTree: []int{-1, 2, 3, -2, 4, -1, -1},
			description:  "包含负数的树",
		},
		{
			name:         "大值树",
			preorder:     []int{100, 50, 25, 75, 150, 125, 175},
			inorder:      []int{25, 50, 75, 100, 125, 150, 175},
			expectedTree: []int{100, 50, 150, 25, 75, 125, 175},
			description:  "大值的BST树",
		},
		{
			name:         "复杂树 - 深度较大",
			preorder:     []int{5, 3, 2, 1, 4, 7, 6, 8, 9},
			inorder:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expectedTree: []int{5, 3, 7, 2, 4, 6, 8, 1, -1, -1, -1, -1, -1, -1, 9},
			description:  "深度较大的树",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 重建树
			result := deduceTree(tt.preorder, tt.inorder)

			// 构建期望的树
			var expected *TreeNode
			if len(tt.expectedTree) > 0 {
				expected = buildTree(tt.expectedTree)
			}

			// 验证树结构是否相同
			if !isSameTree(result, expected) {
				t.Errorf("deduceTree() 重建的树与期望不匹配")
				t.Logf("前序遍历: %v", tt.preorder)
				t.Logf("中序遍历: %v", tt.inorder)
				t.Logf("期望树层序遍历: %v", tt.expectedTree)
			} else {
				// 额外验证：检查重建树的前序和中序遍历是否与输入一致（跳过空树）
				if result != nil || len(tt.preorder) > 0 {
					resultPreorder := getPreorder(result)
					resultInorder := getInorder(result)

					if !reflect.DeepEqual(resultPreorder, tt.preorder) {
						t.Errorf("重建树的前序遍历不匹配: got %v, expected %v", resultPreorder, tt.preorder)
						return
					}
					if !reflect.DeepEqual(resultInorder, tt.inorder) {
						t.Errorf("重建树的中序遍历不匹配: got %v, expected %v", resultInorder, tt.inorder)
						return
					}
				}

				t.Logf("✓ Test passed: %s - %s", tt.name, tt.description)
			}
		})
	}
}

// 测试边界情况
func TestDeduceTreeEdgeCases(t *testing.T) {
	t.Run("空数组", func(t *testing.T) {
		result := deduceTree([]int{}, []int{})
		if result != nil {
			t.Errorf("空数组应该返回 nil, got %v", result)
		}
	})

	t.Run("单节点", func(t *testing.T) {
		result := deduceTree([]int{42}, []int{42})
		if result == nil {
			t.Errorf("单节点应该返回非 nil")
			return
		}
		if result.Val != 42 || result.Left != nil || result.Right != nil {
			t.Errorf("单节点树结构不正确, got %v", result)
		}
	})

	t.Run("验证前序和中序遍历一致性", func(t *testing.T) {
		testCases := []struct {
			preorder []int
			inorder  []int
		}{
			{[]int{1, 2}, []int{2, 1}},
			{[]int{1, 2, 3}, []int{2, 1, 3}},
			{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
		}

		for _, tc := range testCases {
			result := deduceTree(tc.preorder, tc.inorder)
			resultPreorder := getPreorder(result)
			resultInorder := getInorder(result)

			if !reflect.DeepEqual(resultPreorder, tc.preorder) {
				t.Errorf("前序遍历不匹配: input %v, got %v", tc.preorder, resultPreorder)
			}
			if !reflect.DeepEqual(resultInorder, tc.inorder) {
				t.Errorf("中序遍历不匹配: input %v, got %v", tc.inorder, resultInorder)
			}
		}
	})
}

