package binarytree

/*
23、二叉搜索树中的众数
day:2022-6-10
link：https://leetcode.cn/problems/find-mode-in-binary-search-tree/
idea:
如果不是二叉搜索树，最直观的方法一定是把这个树都遍历了，
用map统计频率，把频率排个序，最后取前面高频的元素的集合。

如果是二叉搜索树，中序遍历，然后遍历有序数组的元素出现频率，
从头遍历，那么一定是相邻两个元素作比较，
最后就把出现频率最高的元素输出就可以了。
*/

// 递归中序遍历，边遍历边计数，不适用额外空间。
func findMode(root *TreeNode) []int {
	var res []int
	count := 1
	max := 1
	var pre *TreeNode // 保存上一个节点
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 中序遍历
		traversal(node.Left)
		if pre != nil && pre.Val == node.Val {
			count++
		} else {
			count = 1
		}
		if count >= max { // 这里的逻辑很重要
			if count > max && len(res) > 0 {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			max = count
		}
		pre = node
		traversal(node.Right)
	}
	traversal(root)
	return res
}
