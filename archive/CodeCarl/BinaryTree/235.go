/*
235. 二叉搜索树的最近公共祖先
2023-1-27
link: https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/
question：
	给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
	百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
	例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
answer:
	注意本题的不同：二叉搜索树。
		所有节点的值都是唯一的。
		p、q 为不同节点且均存在于给定的二叉搜索树中。
	当然也可以用普通二叉树的方法解决。

*/
// 普通二叉树版
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如果根节点为空或者为p,q，则他就可能是祖先，返回
	if root == nil || root == p || root == q {
		return root
	}
	// 后序遍历
	// 递归查看左右子树中是否存在p,q
	left := lowestCommonAncestor(root.Left, p, q)   // 左
	right := lowestCommonAncestor(root.Right, p, q) // 右
	// 中
	if left != nil && right != nil { // 两个节点分别在左右子树中
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	// 两个子树中都没有找到节点
	return nil
}

/*
二叉搜索树版：二叉搜索树是有序的。
在有序树里，如果判断一个节点的左子树里有p，右子树里有q呢？
	因为是有序树，所有如果中间节点是 q 和 p 的公共祖先，
	那中节点的数组一定是在[p, q]区间的。
	即 中节点 > p && 中节点 < q 或者 中节点 > q && 中节点 < p。
但是，只要从上到下去遍历，遇到 cur节点是数值在[p, q]区间中则一定可以说明该节点cur就是q 和 p的公共祖先。
 问题来了，不一定是最近公共祖先。
 当我们从上向下去递归遍历，第一次遇到 cur节点是数值在[p, q]区间中，那么cur就是 p和q的最近公共祖先。

不用使用回溯，二叉搜索树自带方向性，可以方便的从上向下查找目标区间，
遇到目标区间内的节点，直接返回。
*/
// 迭代版
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 迭代循环判断
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val { // p,q全部在左子树
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val { // p,q全部在右子树
			root = root.Right
		} else { // p, q分别在俩边子树，则中间节点是祖先
			return root
		}
	}
	// 没有找到
	return nil
}