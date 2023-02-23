package binarytree

/*
20
700、二叉搜索树中的搜索
update::2023-1-26
link:https://leetcode.cn/problems/search-in-a-binary-search-tree/
question:
	给定二叉搜索树（BST）的根节点和一个值。
	你需要在BST中找到节点值等于给定值的节点。
	返回以该节点为根的子树。如果节点不存在，则返回 NULL。
idea:
	二叉搜索树是一个有序树：
	若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
	若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值；
	它的左、右子树也分别为二叉搜索树
	这就决定了，二叉搜索树，递归遍历和迭代遍历和普通二叉树都不一样。
	(普通二叉树的遍历：前序、后序、中序、层次全部遍历，对比就可以了。
	二叉搜索树可以根据特点进行部分搜索，就可以了。)
*/

// 递归实现， 前序遍历寻找节点
func searchBST(root *TreeNode, val int) *TreeNode {
	// 递归结束条件
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val) // return 不能少
	}
	return searchBST(root.Right, val) // return不能少
}

/*
二叉树遍历的迭代法，立刻想起使用栈来模拟深度遍历，使用队列来模拟广度遍历。
对于二叉搜索树可就不一样了，因为二叉搜索树的特殊性，
	也就是节点的有序性，可以不使用辅助栈或者队列就可以写出迭代法。

对于一般二叉树，递归过程中还有回溯的过程，
	例如走一个左方向的分支走到头了，那么要调头，在走右分支。
对于二叉搜索树，不需要回溯的过程，因为节点的有序性就帮我们确定了搜索的方向。
*/
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val > val { // 根节点值大于val，进入左子树搜索
			root = root.Left
		} else if root.Val < val { // 根节点值小于val, 进入右子树搜索
			root = root.Right
		} else {
			return root
		}
	}
	return root
}
