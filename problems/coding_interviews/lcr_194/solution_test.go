package lcr194

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name     string
		treeVals []int
		pVal     int
		qVal     int
		want     int
	}{
		{name: "different subtrees", treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, pVal: 5, qVal: 1, want: 3},
		{name: "same left subtree", treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, pVal: 6, qVal: 2, want: 5},
		{name: "ancestor is one node", treeVals: []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, pVal: 5, qVal: 4, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.treeVals)
			p := findNode(root, tt.pVal)
			q := findNode(root, tt.qVal)
			if p == nil || q == nil {
				t.Fatalf("expected nodes to exist, got p=%v q=%v", p, q)
			}

			got := LowestCommonAncestor(root, p, q)
			if got == nil {
				t.Fatal("LowestCommonAncestor() = nil")
			}
			if got.Val != tt.want {
				t.Fatalf("LowestCommonAncestor() = %d, want %d", got.Val, tt.want)
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

func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}
