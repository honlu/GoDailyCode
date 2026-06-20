/*
437. 路径总和3-中等
2022-11-5
update: 2023-1-23
link: https://leetcode.cn/problems/path-sum-iii/
question:
	给定一个二叉树的根节点root，和一个整数targetSum，
	求该二叉树里节点值之和等于 targetSum 的路径的数目。
	路径不需要从根节点开始，也不需要在叶子节点结束，
	但是路径方向必须是向下的（只能从父节点到子节点）。
answer:
解法一：
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

解法2：层次遍历，迭代实现
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
解法一：深度优先搜索，递归版
*/
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil { // 如果树为空，就返回0条路径
		return 0
	}
	// 否则就返回所有满足条件的路径数量，一共包含三个部分：
	// 当前根节点的，左子树的，右子树的
	return traverse(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

// 深度优先搜索：递归，前序遍历(下边函数算是路径总和1题目的解法略微变形。)
func traverse(node *TreeNode, sum int) int {
	// 当前节点处理
	if node == nil { // 这个判断，也可以挪到递归遍历左右子树时去判断
		return 0
	}
	cnt := 0
	if node.Val == sum {
		cnt++
	}
	// 当前节点的左右子节点递归处理
	cnt += traverse(node.Left, sum-node.Val)
	cnt += traverse(node.Right, sum-node.Val)
	return cnt
}

/*
解法2：层次遍历，迭代实现。
*/
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	// 两个队列和返回结果声明
	res := 0
	queueNode := []*TreeNode{root} // 节点
	queuePath := [][]int{}         // 根节点到当前节点的路径
	queuePath = append(queuePath, []int{root.Val})
	for len(queueNode) > 0 {
		size := len(queueNode)
		for i := 0; i < size; i++ {
			node := queueNode[i]
			path := queuePath[i]
			// 根节点到当前节点的路径，倒过来求和
			val := 0
			for j := len(path) - 1; j >= 0; j-- {
				val += path[j]
				if val == targetSum {
					res++
				}
			}
			if node.Left != nil {
				queueNode = append(queueNode, node.Left)
				// 注意：路径添加
				temp := append([]int{}, path...)
				temp = append(temp, node.Left.Val) // 追加多个元素
				queuePath = append(queuePath, temp)
			}
			if node.Right != nil {
				queueNode = append(queueNode, node.Right)
				temp := append([]int{}, path...) // 追加多个元素
				temp = append(temp, node.Right.Val)
				queuePath = append(queuePath, temp)
			}
		}
		queueNode = queueNode[size:]
		queuePath = queuePath[size:]
	}
	return res
}