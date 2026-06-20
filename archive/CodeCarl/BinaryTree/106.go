package binarytree

/*
17
106.从中序遍历与后序遍历序列构造二叉树
day:2022-6-7
update: 2023-1-24
link: https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
question:
	给定两个整数数组 inorder 和 postorder ，
	其中 inorder 是二叉树的中序遍历，
	postorder 是同一棵树的后序遍历，
	请你构造并返回这颗 二叉树 。

idea:
	关键点：中序遍历的特点，根节点将左右子树的节点一分为二。但并不表示是在最中间，因为左右子树的节点个数可能不一样。
	所以需要从前序或后序中特点找到中间节点。
*/

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 || len(postorder) < 1 || len(inorder) != len(postorder) {
		return nil
	}
	// 先找到根节点(后序遍历的最后一个节点)
	rootVal := postorder[len(postorder)-1]
	// 根节点的索引
	index := findRootIndex(inorder, rootVal)
	// 构造树的根节点
	root := &TreeNode{
		Val:   rootVal,
		Left:  buildTree(inorder[:index], postorder[:index]),                   // 中序和后续遍历的index前面的元素都是左子树的节点值
		Right: buildTree(inorder[index+1:], postorder[index:len(postorder)-1]), // 这里len(postorder)-1是关键
	}

	return root
}

// 根据节点值，找到在中序遍历中的索引
func findRootIndex(inorder []int, target int) int {
	for i, v := range inorder {
		if v == target {
			return i
		}
	}
	return -1
}
