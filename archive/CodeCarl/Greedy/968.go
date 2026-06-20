/*
19
968. 监控二叉树
2022-10-13
2023-2-22 updated by lu
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

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	res := 0 // 安装摄像头的数目

	var backTrack func(root *TreeNode) int // 有返回值表示节点的状态：0表示节点没有被覆盖，1表示当前节点是摄像头，2表示被摄像头覆盖
	backTrack = func(root *TreeNode) int {
		if root == nil { // 空节点默认表示被覆盖
			return 2
		}

		left := backTrack(root.Left)
		right := backTrack(root.Right)

		if left == 2 && right == 2 { // 左右子节点都被覆盖
			return 0 // 父节点无覆盖状态
		} else if left == 0 || right == 0 { // 左右子节点有一个没有被覆盖，父节点就要设置摄像头，使其被覆盖监控
			res++
			return 1
		} else { // 左右子节点有一个安装摄像头
			return 2 // 则当前节点被覆盖
		}

		return -1 // 只是为了返回值，不会走到这里
	}

	if backTrack(root) == 0 { // 如果根节点没有被覆盖
		res++
	}
	return res
}

// 函数提出来写
func minCameraCover(root *TreeNode) int {
	res := 0
	// 情况4-root根节点没有被覆盖
	if backTrack(root, &res) == 0 {
		res++
	}
	return res
}

// 后序遍历：注意有返回值，表示节点的状态
func backTrack(root *TreeNode, res *int) int {
	// 空节点，该节点有覆盖
	if root == nil {
		return 2
	}

	left := backTrack(root.Left, res)   // 先左节点的状态
	right := backTrack(root.Right, res) // 然后右节点的状态
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