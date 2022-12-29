package binarytree

/*
13、相同树
day:2022-6-6
link:https://leetcode.cn/problems/same-tree/
idea:
认清判断对称树本质之后， 对称树的代码 稍作修改 就可以直接用来AC 100.相同的树。
*/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 递归实现
	switch {
	case p == nil && q == nil:
		return true
	case p == nil || q == nil:
		return false
	case p.Val != q.Val:
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
