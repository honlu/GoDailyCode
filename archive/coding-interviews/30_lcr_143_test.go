package main

import (
	"testing"
)

func TestIsSubStructure(t *testing.T) {
	tests := []struct {
		name     string
		aVals    []int // 树A的层序遍历数组，-1表示空节点
		bVals    []int // 树B的层序遍历数组，-1表示空节点
		expected bool
	}{
		{
			name:     "空树情况 - A和B都为空",
			aVals:    []int{},
			bVals:    []int{},
			expected: false,
		},
		{
			name:     "空树情况 - A为空，B不为空",
			aVals:    []int{},
			bVals:    []int{1},
			expected: false,
		},
		{
			name:     "空树情况 - A不为空，B为空",
			aVals:    []int{1},
			bVals:    []int{},
			expected: false,
		},
		{
			name:     "B是A的完整子树",
			aVals:    []int{3, 4, 5, 1, 2},
			bVals:    []int{4, 1},
			expected: true,
		},
		{
			name:     "B不是A的子结构",
			aVals:    []int{3, 4, 5, 1, 2},
			bVals:    []int{4, 3},
			expected: false,
		},
		{
			name:     "B等于A",
			aVals:    []int{1, 2, 3},
			bVals:    []int{1, 2, 3},
			expected: true,
		},
		{
			name:     "B在A的左子树中",
			aVals:    []int{3, 4, 5, 1, 2, -1, -1},
			bVals:    []int{4, 1, 2},
			expected: true,
		},
		{
			name:     "B在A的右子树中",
			aVals:    []int{3, 4, 5, -1, -1, 1, 2},
			bVals:    []int{5, 1, 2},
			expected: true,
		},
		{
			name:     "B是A的一部分但结构不匹配",
			aVals:    []int{3, 4, 5, 1, 2},
			bVals:    []int{4, 1, 3},
			expected: false,
		},
		{
			name:     "单节点 - B是A",
			aVals:    []int{1},
			bVals:    []int{1},
			expected: true,
		},
		{
			name:     "单节点 - B不是A",
			aVals:    []int{1},
			bVals:    []int{2},
			expected: false,
		},
		{
			name:     "复杂情况 - B在A的右子树深层中",
			aVals:    []int{1, 2, 3, 4, 5, 6, 7},
			bVals:    []int{3, 6, 7},
			expected: true,
		},
		{
			name:     "B的值在A中但结构不完整",
			aVals:    []int{10, 12, 6, 8, 3, 11},
			bVals:    []int{10, 12, 6, 8},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			A := buildTree(tt.aVals)
			B := buildTree(tt.bVals)
			result := isSubStructure(A, B)

			if result != tt.expected {
				t.Errorf("isSubStructure() = %v, expected %v", result, tt.expected)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}
