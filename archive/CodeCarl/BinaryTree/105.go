/*
105. 从前序与中序遍历序列构造二叉树
2023-1-24
link: https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
question:
	给定两个整数数组 preorder 和 inorder ，
	其中 preorder 是二叉树的先序遍历，
	inorder 是同一棵树的中序遍历，
	请构造二叉树并返回其根节点。

answer:
*/
// 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) < 1 || len(preorder) < 1 {
		return nil
	}
	// 先找到根节点（前序遍历的第一个就是根节点）
	rootValue := preorder[0]
	// 从中序遍历中找到一分为二的点
	index := findRootIndex(inorder, rootValue)
	// 构造root
	root := &TreeNode{
		Val:   rootValue,                                       // 下面递归是关键
		Left:  buildTree(preorder[1:index+1], inorder[:index]), // //将后序遍历一分为二，左边为左子树，右边为右子树
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}
	return root
}

// 从中序遍历中找到一分为二的点，左边为左子树，右边右子树
func findRootIndex(inorder []int, target int) int {
	for i, v := range inorder {
		if v == inorder[i] {
			return i
		}
	}
	return -1
}