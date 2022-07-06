package binarytree

/*
24、二叉树的最近公共祖先
day: 2022-6-10
link: https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
idea:
首先想的是要是能自底向上查找就好了，这样就可以找到公共祖先了。
那么二叉树如何可以自底向上查找呢？
回溯啊，二叉树回溯的过程就是从低到上。
后序遍历就是天然的回溯过程，最先处理的一定是叶子节点。
接下来就看如何判断一个节点是节点q和节点p的公共公共祖先呢。

首先最容易想到的一个情况：如果找到一个节点，发现左子树出现结点p，
右子树出现节点q，或者 左子树出现结点q，右子树出现节点p，
那么该节点就是节点p和q的最近公共祖先。
但是很多人容易忽略一个情况，就是节点本身p(q)，它拥有一个子孙节点q(p)。

使用后序遍历，回溯的过程，就是从低向上遍历节点，一旦发现满足第一种情况的节点，就是最近公共节点了。

但是如果p或者q本身就是最近公共祖先呢？其实只需要找到一个节点是p或者q的时候，直接返回当前节点，
无需继续递归子树。
如果接下来的遍历中找到了后继节点满足第一种情况则修改返回值为后继节点，
否则，继续返回已找到的节点即可。
为什么满足第一种情况的节点一定是p或q的后继节点呢?
大家可以仔细思考一下。
*/

// 后序遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q { // 递归结束条件
		return root
	}
	// 后序遍历
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 左右两边都不为空，则根节点为祖先
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
