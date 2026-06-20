package main

// LCR 124. 推理二叉树
func deduceTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 树的先序遍历和中序遍历，复原原有树结构
	item := preorder[0]
	root := &TreeNode{
		Val: item,
	}
	var leftPreorder, leftInorder []int
	var rightPreorder, rightInorder []int
	for i, num := range inorder {
		if num == item {
			leftPreorder = preorder[1 : 1+i] // 注意边界
			leftInorder = inorder[:i]
			rightPreorder = preorder[1+i:] // 注意边界
			rightInorder = inorder[i+1:]
		}
	}
	if len(leftPreorder) > 0 { // 注意先判断
		root.Left = deduceTree(leftPreorder, leftInorder)
	}
	if len(rightPreorder) > 0 {
		root.Right = deduceTree(rightPreorder, rightInorder)
	}

	return root
}
