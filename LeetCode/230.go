/*
230. 二叉搜索树中第K个最小值
2023-3-6 by lu
link: https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
question:
	给定一个二叉搜索树的根节点 root ，和一个整数 k ，
	 请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
answer:
	二叉搜索树的中序遍历是有序的。
	所以可以通过中序遍历思想来实现！
*/
// 保存所有节点方式，返回第k个最小值
func kthSmallest(root *TreeNode, k int) int {
	// 二叉搜索树中序遍历是有序的
	var backTrack func(root *TreeNode, temp *[]int)
	backTrack = func(root *TreeNode, temp *[]int) {
		if root == nil {
			return
		}
		backTrack(root.Left, temp)
		*temp = append(*temp, root.Val)
		backTrack(root.Right, temp)
	}
	var temp []int
	backTrack(root, &temp)
	return temp[k-1]
}

// 利用k标志，直接返回符合的值
func kthSmallest(root *TreeNode, k int) int {
	// 二叉搜索树中序遍历是有序的
	var backTrack func(root *TreeNode)
	var res int
	backTrack = func(root *TreeNode) {
		if root == nil {
			return
		}
		backTrack(root.Left)
		k-- // 注意，必须使用全局k；如果当作参数传入很容易覆盖！
		if k == 0 {
			res = root.Val
			return
		}
		backTrack(root.Right)
	}
	backTrack(root)
	return res
}