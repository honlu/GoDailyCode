/*
剑指offer26. 树的子结构
2023-3-17
link: https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/
question:
	输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
	B是A的子结构， 即 A中有出现和B相同的结构和节点值。
answer:
	首先找到A中节点和B根节点相同的节点，然后判断相应子结构是否相同
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil && B == nil {
		return true
	} else if A == nil || B == nil {
		return false
	}
	// 遍历A树
	var traversal func(a *TreeNode, b *TreeNode)
	var res bool
	traversal = func(a *TreeNode, b *TreeNode) {
		// 前序遍历
		if a == nil {
			return
		}
		if a.Val == b.Val {
			res = isSub(a, b)
		}
		traversal(a.Left, b)
		traversal(a.Right, b)
	}
	traversal(A, B)
	return res
}

// 判断子结构是否相同
func isSub(a *TreeNode, b *TreeNode) bool {
	if (a == nil && b == nil) || (a != nil && b == nil) { // 要注意a可能有些节点不为空， b节点为空也是对应的
		return true
	} else if a == nil && b != nil {
		return false
	}
	if a.Val == b.Val {
		return isSub(a.Left, b.Left) && isSub(a.Right, b.Right)
	} else {
		return false
	}
}