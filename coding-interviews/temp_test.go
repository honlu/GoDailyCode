package main

import "fmt"

func main() {
	treeVals := []int{1, 2, 3, -1, 4, 5, 6, -1, -1, 7, 8}
	root := buildTree(treeVals)

	var printTree func(*TreeNode, int)
	printTree = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		for i := 0; i < depth; i++ {
			fmt.Print("  ")
		}
		fmt.Printf("%d\n", node.Val)
		printTree(node.Left, depth+1)
		printTree(node.Right, depth+1)
	}

	fmt.Println("树结构:")
	printTree(root, 0)

	p := findNodeInBT(root, 7)
	q := findNodeInBT(root, 8)

	if p != nil && q != nil {
		fmt.Printf("\n节点7的值: %d\n", p.Val)
		fmt.Printf("节点8的值: %d\n", q.Val)
		result := lowestCommonAncestor194(root, p, q)
		if result != nil {
			fmt.Printf("LCA: %d\n", result.Val)
		}
	}
}
