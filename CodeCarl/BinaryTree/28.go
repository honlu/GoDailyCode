package binarytree

/*
28. 将有序数组转换为·高度平衡·二叉搜索树
day:2022-6-15
link:https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
idea:
这里不用强调平衡二叉搜索树，数组构造二叉树，构成平衡树是自然而然的事情，因为大家默认都是从数组中间位置取值作为节点元素，
一般不会随机取，所以想构成不平衡的二叉树是自找麻烦。
本质就是寻找分割点，分割点作为当前节点，然后递归左区间和右区间。
*/

// 递归实现
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 { // 递归终止条件，最后数组为空则可以返回
		return nil
	}
	root := &TreeNode{
		Val:   nums[len(nums)/2], // 按照二叉搜索树的特点，从中间构造根节点
		Left:  nil,
		Right: nil,
	}
	root.Left = sortedArrayToBST(nums[:len(nums)/2])    // 数组的左半部分作为左子树
	root.Right = sortedArrayToBST(nums[len(nums)/2+1:]) // 数组的右半部分作为右子树
	return root
}
