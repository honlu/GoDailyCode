package binarytree

/*
10
222、完全二叉树的节点个数
day:2022-6-5
update: 2023-1-18
question:
	给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
	完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，
	其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。
	若最底层为第 h 层，则该层包含 1~ 2h 个节点。

idea:
	分外普通二叉树解法和利用完全二叉树特性的递归解法.
	普通二叉树解法：
		直接层次遍历，或者无脑遍历就可以！
	完全二叉树特性：
*/

// 普通二叉树，求有多少节点，直接无脑遍历相加算长度就行
// 递归版本
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Right != nil {
		res += countNodes(root.Right)
	}
	if root.Left != nil {
		res += countNodes(root.Left)
	}
	return res
}

// 迭代版本：层次遍历
func countNodes(root *TreeNode) int {
	// 层次遍历解决:每层节点数相加就可以
	if root == nil {
		return 0
	}
	// 队列初始化
	queue := []*TreeNode{root}
	res := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res += size          // 这里更新节点数
		queue = queue[size:] // 这里更新每层的节点
	}
	return res
}

// 利用完全二叉树特性的递归解法
func countNodesSpecial(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := 0, 0
	// 直接通过最左侧的树节点和最右侧的树节点判断是否是满二叉树
	left, right := root.Left, root.Right
	for left != nil {
		left = left.Left
		leftHeight += 1
	}
	for right != nil {
		right = right.Right
		rightHeight += 1
	}
	if leftHeight == rightHeight {
		return (2 << leftHeight) - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1 // 注意这个递归，如果某个节点左右子树不是满二叉树，就会+1，并且一直下去
}
