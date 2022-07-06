package binarytree

/*
17-从中序遍历与后序遍历序列构造二叉树
day:2022-6-7
link: https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
从前序与中序遍历序列构造二叉树
link:https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
idea:
*/

// 从中序遍历中找到一分为二的点，左边为左子树，右边右子树
func findRootIndex(inorder []int, target int) int {
	for i := 0; i < len(inorder); i++ {
		if target == inorder[i] {
			return i
		}
	}
	return -1
}

// 从中序遍历与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 {
		return nil
	}
	// 先找到根节点（后序遍历的最后一个就是根节点）
	nodeValue := postorder[len(postorder)-1]
	// 从中序遍历中找到一分为二的点
	left := findRootIndex(inorder, nodeValue)
	// 构造root
	root := &TreeNode{
		Val:   nodeValue,
		Left:  buildTree(inorder[:left], postorder[:left]), // //将后序遍历一分为二，左边为左子树，右边为右子树
		Right: buildTree(inorder[left+1:], postorder[left:len(postorder)-1]),
	}
	return root
}

// 从前序与中序遍历序列构造二叉树
func buildTreePre(preorder []int, inorder []int) *TreeNode {
	if len(inorder) < 1 || len(preorder) < 1 {
		return nil
	}
	// 先找到根节点（前序遍历的第一个就是根节点）
	nodeValue := preorder[0]
	// 从中序遍历中找到一分为二的点
	left := findRootIndex(inorder, nodeValue)
	// 构造root
	root := &TreeNode{
		Val:   nodeValue,
		Left:  buildTree(preorder[1:left+1], inorder[:left]), // //将后序遍历一分为二，左边为左子树，右边为右子树
		Right: buildTree(preorder[left+1:], inorder[left+1:]),
	}
	return root
}
