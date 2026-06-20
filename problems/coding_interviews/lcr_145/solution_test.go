package lcr145

import "testing"

func TestCheckSymmetricTree(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{name: "empty tree", input: []int{}, want: true},
		{name: "symmetric tree", input: []int{1, 2, 2, 3, 4, 4, 3}, want: true},
		{name: "value mismatch", input: []int{1, 2, 3}, want: false},
		{name: "structure mismatch", input: []int{1, 2, 2, -1, 3, -1, 3}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckSymmetricTree(buildTree(tt.input)); got != tt.want {
				t.Fatalf("CheckSymmetricTree() = %v, want %v", got, tt.want)
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
