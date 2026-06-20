package main

import "testing"

func TestLowestCommonAncestor194DeepDebugScenario(t *testing.T) {
	treeVals := []int{1, 2, 3, -1, 4, 5, 6, -1, -1, 7, 8}
	root := buildTree(treeVals)

	p := findNodeInBT(root, 7)
	q := findNodeInBT(root, 8)
	if p == nil || q == nil {
		t.Fatalf("expected both nodes to exist, got p=%v q=%v", p, q)
	}

	result := lowestCommonAncestor194(root, p, q)
	if result == nil {
		t.Fatal("expected LCA node, got nil")
	}
	if result.Val != 4 {
		t.Fatalf("expected LCA value 4, got %d", result.Val)
	}
}
