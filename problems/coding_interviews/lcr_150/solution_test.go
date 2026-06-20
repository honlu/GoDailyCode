package lcr150

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  [][]int
	}{
		{name: "empty tree", input: []int{}, want: nil},
		{name: "single node", input: []int{1}, want: [][]int{{1}}},
		{name: "three levels", input: []int{3, 9, 20, -1, -1, 15, 7}, want: [][]int{{3}, {9, 20}, {15, 7}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LevelOrder(buildTree(tt.input)); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("LevelOrder() = %v, want %v", got, tt.want)
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
