package main

import (
	"testing"
)

// isSameTree 比较两棵树是否相同
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func TestCodec(t *testing.T) {
	tests := []struct {
		name     string
		treeVals []int // 层序遍历数组，-1表示空节点
	}{
		{
			name:     "空树",
			treeVals: []int{},
		},
		{
			name:     "单节点",
			treeVals: []int{1},
		},
		{
			name:     "简单树 - 只有左子树",
			treeVals: []int{1, 2, -1},
		},
		{
			name:     "简单树 - 只有右子树",
			treeVals: []int{1, -1, 3},
		},
		{
			name:     "简单二叉树",
			treeVals: []int{1, 2, 3},
		},
		{
			name:     "完整二叉树",
			treeVals: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "不平衡树",
			treeVals: []int{1, 2, 3, 4, -1, -1, 5},
		},
		{
			name:     "复杂树 - 包含负数",
			treeVals: []int{-1, 2, 3, -2, 4},
		},
		{
			name:     "复杂树 - 大值",
			treeVals: []int{100, 50, 150, 25, 75, 125, 175},
		},
		{
			name:     "左斜树",
			treeVals: []int{1, 2, -1, 3},
		},
		{
			name:     "右斜树",
			treeVals: []int{1, -1, 2, -1, -1, -1, 3},
		},
		{
			name:     "复杂树 - 多层级",
			treeVals: []int{5, 3, 7, 2, 4, 6, 8, 1, -1, -1, -1, -1, -1, -1, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 构建原始树
			originalTree := buildTree(tt.treeVals)

			// 创建 Codec 实例
			ser := Constructor()
			deser := Constructor()

			// 序列化
			data := ser.serialize(originalTree)
			t.Logf("序列化结果: %s", data)

			// 反序列化
			deserializedTree := deser.deserialize(data)

			// 验证反序列化后的树是否与原树相同
			if !isSameTree(originalTree, deserializedTree) {
				t.Errorf("反序列化后的树与原树不匹配")
				t.Logf("原始树层序遍历: %v", tt.treeVals)
				// 可以添加更详细的错误信息
			} else {
				t.Logf("✓ Test passed: %s", tt.name)
			}
		})
	}
}

// 测试序列化和反序列化的边界情况
func TestCodecEdgeCases(t *testing.T) {
	t.Run("空字符串反序列化", func(t *testing.T) {
		deser := Constructor()
		result := deser.deserialize("")
		if result != nil {
			t.Errorf("空字符串反序列化应该返回 nil, got %v", result)
		}
	})

	t.Run("只有nil的字符串", func(t *testing.T) {
		ser := Constructor()
		deser := Constructor()

		// 序列化空树
		data := ser.serialize(nil)
		if data != "nil" {
			t.Errorf("空树序列化应该返回 'nil', got %s", data)
		}

		// 反序列化
		result := deser.deserialize(data)
		if result != nil {
			t.Errorf("反序列化 'nil' 应该返回 nil, got %v", result)
		}
	})

	t.Run("序列化后再反序列化空树", func(t *testing.T) {
		ser := Constructor()
		deser := Constructor()

		data := ser.serialize(nil)
		result := deser.deserialize(data)

		if result != nil {
			t.Errorf("空树序列化后再反序列化应该返回 nil, got %v", result)
		}
	})
}
