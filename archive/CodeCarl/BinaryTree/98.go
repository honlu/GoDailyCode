package binarytree

import "math"

/*
21
98.验证二叉搜索树
day:2022-6-9
update: 2023-1-26
link:https://leetcode.cn/problems/validate-binary-search-tree/
question:
	给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
	有效 二叉搜索树定义如下：
		节点的左子树只包含 小于 当前节点的数。
		节点的右子树只包含 大于 当前节点的数。
		所有左子树和右子树自身必须也是二叉搜索树。
idea:
	注意上面的二叉搜索树的定义。
	要知道中序遍历下，输出的二叉搜索树节点的数值是有序序列。
	有了这个特性，验证二叉搜索树，就相当于变成了判断一个序列是不是递增的了。

	注意两个坑：
	1.不能单纯的比较左节点小于中间节点，右节点大于中间节点就完事了
		因为要比较的是 左子树所有节点小于中间节点，右子树所有节点大于中间节点。
	2.最小节点 可能是int的最小值，如果这样使用最小的int来比较也是不行的。

*/

// 递归中序遍历解法，利用二叉搜索树的中序遍历是有序的，
/// 所以设置一个上一个节点，就可以依次判断了
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode // 记录前一个节点
	var traversal func(node *TreeNode) bool
	// 递归函数定义
	traversal = func(node *TreeNode) bool {
		// 递归结束条件
		if node == nil { // 二叉搜索树也可以是空树
			return true
		}
		// 单层逻辑代码
		// 左
		left := traversal(node.Left) // 注意这里全局变量pre也会进入
		// 中
		if pre != nil && pre.Val >= node.Val {
			return false
		}
		pre = node
		// 右
		right := traversal(node.Right)
		return left && right
	}
	return traversal(root)
}

/*
迭代中序遍历，稍作改动.
中序遍历是二叉树的一种遍历方式，它先遍历左子树，再遍历根节点，
最后遍历右子树。而我们二叉搜索树保证了左子树的节点的值均小于根节点的值，
根节点的值均小于右子树的值，因此中序遍历以后得到的序列一定是升序序列。

启示我们在中序遍历的时候实时检查当前节点的值是否大于前一个中序遍历到的节点的值即可。
如果均大于说明这个序列是升序的，整棵树是二叉搜索树，否则不是.

*/
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 栈初始化
	stack := []*TreeNode{}
	inorder := math.MinInt64 // 中间节点最小值
	// 中序遍历
	for len(stack) > 0 || root != nil {
		// 进栈
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 出栈
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}
