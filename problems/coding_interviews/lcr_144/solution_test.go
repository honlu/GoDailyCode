package lcr144

import (
	"reflect"
	"testing"
)

func TestFlipTree(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{name: "empty tree", input: []int{}, want: []int{}},
		{name: "simple tree", input: []int{1, 2, 3}, want: []int{1, 3, 2}},
		{name: "complete tree", input: []int{4, 2, 7, 1, 3, 6, 9}, want: []int{4, 7, 2, 9, 6, 3, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := treeToLevelOrder(FlipTree(buildTree(tt.input)))
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("FlipTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		if i < len(vals) {
			if vals[i] != -1 {
				node.Left = &TreeNode{Val: vals[i]}
				queue = append(queue, node.Left)
			} else {
				queue = append(queue, nil)
			}
			i++
		}
		if i < len(vals) {
			if vals[i] != -1 {
				node.Right = &TreeNode{Val: vals[i]}
				queue = append(queue, node.Right)
			} else {
				queue = append(queue, nil)
			}
			i++
		}
	}
	return root
}

func treeToLevelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, -1)
			continue
		}
		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	for len(result) > 0 && result[len(result)-1] == -1 {
		result = result[:len(result)-1]
	}
	return result
}
