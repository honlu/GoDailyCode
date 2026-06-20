package main

import (
	"reflect"
	"testing"
)

func TestPathTarget(t *testing.T) {
	tests := []struct {
		name     string
		input    []int // 树的层序遍历数组，-1表示空节点
		target   int
		expected [][]int
	}{
		{
			name:     "空树",
			input:    []int{},
			target:   0,
			expected: [][]int{},
		},
		{
			name:     "单节点树 - 匹配目标",
			input:    []int{5},
			target:   5,
			expected: [][]int{{5}},
		},
		{
			name:     "单节点树 - 不匹配目标",
			input:    []int{5},
			target:   3,
			expected: [][]int{},
		},
		{
			name:     "简单树 - 多条路径匹配目标22",
			input:    []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, 5, 1},
			target:   22,
			expected: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			name:     "简单树 - 一条路径匹配目标26",
			input:    []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, 5, 1},
			target:   26,
			expected: [][]int{{5, 8, 13}},
		},
		{
			name:     "简单树 - 没有路径匹配",
			input:    []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, 5, 1},
			target:   100,
			expected: [][]int{},
		},
		{
			name:     "简单二叉树 - 根到左叶子",
			input:    []int{1, 2, 3},
			target:   3,
			expected: [][]int{{1, 2}},
		},
		{
			name:     "简单二叉树 - 根到右叶子",
			input:    []int{1, 2, 3},
			target:   4,
			expected: [][]int{{1, 3}},
		},
		{
			name:     "简单二叉树 - 两条路径都匹配",
			input:    []int{1, 2, 3},
			target:   3,
			expected: [][]int{{1, 2}},
		},
		{
			name:     "复杂树 - 多条路径",
			input:    []int{10, 5, -3, 3, 2, -1, 11, 3, -2, -1, 1},
			target:   18,
			expected: [][]int{{10, 5, 2, 1}, {10, -3, 11}},
		},
		{
			name:     "负数节点",
			input:    []int{-2, -1, -3},
			target:   -3,
			expected: [][]int{},
		},
		{
			name:     "目标值为0",
			input:    []int{1, -1, 2, -1, -1, -1, -1},
			target:   0,
			expected: [][]int{},
		},
		{
			name:     "目标值为0且路径和为0",
			input:    []int{0},
			target:   0,
			expected: [][]int{{0}},
		},
		{
			name:     "左斜树 - 匹配",
			input:    []int{1, 2, -1, 3},
			target:   6,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name:     "右斜树 - 匹配",
			input:    []int{1, -1, 2, -1, -1, -1, 3},
			target:   6,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name:     "完全二叉树 - 多条路径",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			target:   7,
			expected: [][]int{{1, 2, 4}},
		},
		{
			name:     "完全二叉树 - 目标值10",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			target:   10,
			expected: [][]int{{1, 3, 6}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.input)
			result := pathTarget(root, tt.target)

			// 规范化空 slice（nil 和 [] 在 reflect.DeepEqual 中不相等）
			if result == nil {
				result = [][]int{}
			}
			expected := tt.expected
			if expected == nil {
				expected = [][]int{}
			}

			if !reflect.DeepEqual(result, expected) {
				t.Errorf("pathTarget() = %v, expected %v", result, expected)
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

