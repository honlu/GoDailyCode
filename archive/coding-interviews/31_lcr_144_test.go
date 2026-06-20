package main

import (
	"reflect"
	"testing"
)

// treeToLevelOrder 将二叉树转换为层序遍历数组（用于测试验证）
// -1 表示空节点
func treeToLevelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, -1)
			continue
		}

		result = append(result, node.Val)
		queue = append(queue, node.Left)
		queue = append(queue, node.Right)
	}

	// 移除末尾的 -1（空节点）
	for len(result) > 0 && result[len(result)-1] == -1 {
		result = result[:len(result)-1]
	}

	return result
}

func TestFlipTree(t *testing.T) {
	tests := []struct {
		name     string
		input    []int // 原始树的层序遍历数组，-1表示空节点
		expected []int // 翻转后树的层序遍历数组，-1表示空节点
	}{
		{
			name:     "空树",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "单节点树",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "简单二叉树 - 两个子节点",
			input:    []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "完全二叉树",
			input:    []int{4, 2, 7, 1, 3, 6, 9},
			expected: []int{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name:     "不完全二叉树 - 左子树为空",
			input:    []int{1, -1, 2, -1, -1, 3},
			expected: []int{1, 2, -1, -1, 3},
		},
		{
			name:     "不完全二叉树 - 右子树为空",
			input:    []int{1, 2, -1, 3},
			expected: []int{1, -1, 2, -1, 3},
		},
		{
			name:     "复杂二叉树",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{1, 3, 2, 7, 6, 5, 4},
		},
		{
			name:     "左斜树",
			input:    []int{1, 2, -1, 3},
			expected: []int{1, -1, 2, -1, 3},
		},
		{
			name:     "右斜树",
			input:    []int{1, -1, 2, -1, -1, -1, 3},
			expected: []int{1, 2, -1, 3},
		},
		{
			name:     "多层复杂树",
			input:    []int{3, 9, 20, -1, -1, 15, 7},
			expected: []int{3, 20, 9, 7, 15},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			result := flipTree(root)
			resultVals := treeToLevelOrder(result)

			if !reflect.DeepEqual(resultVals, tt.expected) {
				t.Errorf("flipTree() = %v, expected %v", resultVals, tt.expected)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

// 测试翻转两次应该回到原始树
func TestFlipTreeTwice(t *testing.T) {
	tests := []struct {
		name  string
		input []int
	}{
		{
			name:  "单节点树",
			input: []int{1},
		},
		{
			name:  "完全二叉树",
			input: []int{4, 2, 7, 1, 3, 6, 9},
		},
		{
			name:  "复杂二叉树",
			input: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:  "不完全二叉树",
			input: []int{3, 9, 20, -1, -1, 15, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			originalVals := treeToLevelOrder(root)

			// 翻转一次
			flipped := flipTree(root)
			flippedVals := treeToLevelOrder(flipped)

			// 翻转第二次，应该回到原始状态
			flippedTwice := flipTree(flipped)
			flippedTwiceVals := treeToLevelOrder(flippedTwice)

			if !reflect.DeepEqual(flippedTwiceVals, originalVals) {
				t.Errorf("flipTree(flipTree()) = %v, expected original %v", flippedTwiceVals, originalVals)
				t.Errorf("First flip result: %v", flippedVals)
			} else {
				t.Logf("✓ Test passed: %s (flip twice returns to original)", tt.name)
			}
		})
	}
}

