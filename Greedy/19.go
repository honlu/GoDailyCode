/*
19、监控二叉树
2022-10-13
link:968-https://leetcode.cn/problems/binary-tree-cameras/
question:
	给定一个二叉树，我们在树的节点上安装摄像头。
	节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
	计算监控树的所有节点所需的最小摄像头数量。
answer:
	树从下往上看，局部最优：让叶子节点的父节点安摄像头，所用摄像头最少，
	整体最优：全部摄像头数量所用最少！
	局部最优推出全局最优，找不出反例，那么就按照贪心来！
	大体思路：
		从低到上，先给叶子节点父节点放个摄像头，
		然后隔两个节点放一个摄像头，直至到二叉树头结点。
重点：后序遍历可以实现从低向上推到。
	判断节点的状态：0-该节点无覆盖 1-本节点有摄像头 2-本节点有覆盖
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minCameraCover(root *TreeNode) int {
	res := 0
	// 情况4-root根节点没有被覆盖
	if backTrack(root, &res) == 0 {
		res++
	}
	return res
}

// 后序遍历
func backTrack(root *TreeNode, res *int) int {
	// 空节点，该节点有覆盖
	if root == nil {
		return 2
	}

	left := backTrack(root.Left)   // 先左节点
	right := backTrack(root.Right) // 然后右节点
	// 最后中间节点-后序遍历
	// 情况1-左右节点都有覆盖
	if left == 2 && right == 2 {
		return 0
	}

	// 情况2- 左右节点有一个没有被覆盖。则当前节点需要加摄像头
	if left == 0 || right == 0 {
		*res++
		return 1
	}

	// 情况3-左右节点存在一个摄像头，则当前节点会被覆盖
	if left == 1 || right == 1 {
		return 2
	}
	// 走不到的逻辑地方，保证返回语法正确
	return -1
}