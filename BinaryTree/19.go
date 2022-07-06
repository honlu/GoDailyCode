package binarytree

/*
19、合并二叉树
day:2022-6-9
link:https://leetcode.cn/problems/merge-two-binary-trees/
idea:合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，
否则不为 NULL 的节点将直接作为新二叉树的节点。
其实和遍历一个树逻辑是一样的，只不过传入两个树的节点，同时操作。
哪种遍历都是可以的。
*/

// 前序遍历-简洁版，以第一个二叉树为改动
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2 // 注意这里返回的节点
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

// 迭代版的话，需要队列或栈来存储当前节点
