/*
437. 路径总和3-中等
2022-11-5
link: https://leetcode.cn/problems/path-sum-iii/
question:
	给定一个二叉树的根节点root，和一个整数targetSum，
	求该二叉树里节点值之和等于 targetSum 的路径的数目。
	路径不需要从根节点开始，也不需要在叶子节点结束，
	但是路径方向必须是向下的（只能从父节点到子节点）。
answer:
	深度优先搜索算法！递归+回溯方法
	1、确定递归函数的参数和返回类型：
		参数：需要二叉树的根节点，还需要一个计数器：用来计算二叉树的一条边之和是否正好是目标和，计数器为int型。
		返回值：返回当前根节点出发路径和等于给定值的路径数量
	2、确定终止条件：
		首先计数器如何统计这一条路径的和呢？
			采用递减的方式，让计数器count初始为目标和，然后减去遍历路径节点上的数值。
			如果最后count==0,同时找到叶子节点的话，说明找到了目标和。
			如果遍历到叶子节点，count!=0,就是没找到。
		注意：不要去累加然后判断是否等于目标和，这样代码比较麻烦。
	3、确定单层递归的逻辑：[重点]
	https://leetcode.cn/problems/path-sum-iii/solution/custerxue-xi-bi-ji-er-cha-shu-de-di-gui-he-dfs-by-/
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil { // 如果树为空，就返回0条路径
		return 0
	}
	// 否则就返回所有满足条件的路径数量，一共包含三个部分：
	// 当前根节点的，左子树的，右子树的
	return traverse(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func traverse(node *TreeNode, sum int) int {
	if node == nil { // 这个判断，也可以挪到递归遍历左右子树时去判断
		return 0
	}
	cnt := 0
	if node.Val == sum {
		cnt++
	}

	cnt += traverse(node.Left, sum-node.Val)
	cnt += traverse(node.Right, sum-node.Val)
	return cnt
}