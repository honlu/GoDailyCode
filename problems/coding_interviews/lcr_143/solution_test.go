package lcr143

import "testing"

func TestIsSubStructure(t *testing.T) {
	tests := []struct {
		name  string
		aVals []int
		bVals []int
		want  bool
	}{
		{name: "empty trees", aVals: []int{}, bVals: []int{}, want: false},
		{name: "sub structure exists", aVals: []int{3, 4, 5, 1, 2}, bVals: []int{4, 1}, want: true},
		{name: "structure mismatch", aVals: []int{3, 4, 5, 1, 2}, bVals: []int{4, 3}, want: false},
		{name: "same tree", aVals: []int{1, 2, 3}, bVals: []int{1, 2, 3}, want: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubStructure(buildTree(tt.aVals), buildTree(tt.bVals)); got != tt.want {
				t.Fatalf("IsSubStructure() = %v, want %v", got, tt.want)
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
