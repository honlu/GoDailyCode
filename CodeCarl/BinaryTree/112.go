package binarytree

/*
16
112、路径总和
day:2022-6-7
link:https://leetcode.cn/problems/path-sum/
question:
	给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。
	判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。
	如果存在，返回 true ；否则，返回 false 。

	叶子节点 是指没有子节点的节点。

idea:
	要遍历从根节点到叶子节点的的路径看看总和是不是目标和。
	方式一：
		深度优先搜索：前序遍历+回溯
		递归版类似前序遍历
	方式二：
		广度优先搜索：层次遍历
		使用广度优先搜索的方式，记录从根节点到当前节点的路径和，
		以防止重复计算。
		所以使用两个队列，分别存储将要遍历的节点，
		以及根节点到这些节点的路径和即可。

*/

// 递归实现路经总和：1还比较简单，只是判断是否存在
// （递归版：类似前序遍历）
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 中间节点显出
	targetSum -= root.Val                                        // 将targetSum在遍历每层的时候都减去本层节点的值
	if root.Left == nil && root.Right == nil && targetSum == 0 { // 如果剩余的targetSum为0，则正好是符合的结果
		return true
	}
	// 然后左右节点递归
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum) // 否则递归继续找。
}

// 递归实现：返回所有存在的路径.
// 深度优先搜索（前中序遍历都可以，这里采用前序）：
// 枚举每一条从根节点到叶子节点的路径。
// 当我们遍历到叶子节点，且此时路径和恰为目标和时，我们就找到了一条满足条件的路径。
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	path := []int{}
	var dfs func(node *TreeNode, left int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return // 回溯
		}
		left -= node.Val              // 遍历每层的时候，目标值减去当前节点
		path = append(path, node.Val) // 将当前节点放到路径记录中
		if node.Left == nil && node.Right == nil && left == 0 {
			// 当前节点为叶子节点，并且满足和。则添加路径，已经回溯
			res = append(res, append([]int(nil), path...)) // 在空的切片多添加多个元素，再添加到结果集中
			return                                         // 回溯
		}
		if node.Left != nil {
			dfs(node.Left, left)
			path = path[:len(path)-1] // 回溯
		}
		if node.Right != nil {
			dfs(node.Right, left)
			path = path[:len(path)-1] // 回溯
		}
	}
	dfs(root, targetSum)
	return res
}

/*
层次遍历
时间复杂度：O(N)，其中N 是树的节点数。对每个节点访问一次。
空间复杂度：O(N)，其中N 是树的节点数。
	空间复杂度主要取决于队列的开销，队列中的元素个数不会超过树的节点数。
*/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 两个队列初始化
	queueNode := []*TreeNode{root} // 要遍历的节点
	queueVal := []int{root.Val}    // 根节点到当前节点的总和
	// 层次遍历开始
	for len(queueNode) > 0 {
		size := len(queueNode)
		for i := 0; i < size; i++ {
			node, val := queueNode[i], queueVal[i]
			if node.Left == nil && node.Right == nil { // 叶子节点
				if val == targetSum { // 判断根节点到当前叶子节点的路径和是否满足要求
					return true
				}
			}
			if node.Left != nil {
				queueNode = append(queueNode, node.Left)
				queueVal = append(queueVal, val+node.Left.Val)
			}
			if node.Right != nil {
				queueNode = append(queueNode, node.Right)
				queueVal = append(queueVal, val+node.Right.Val)
			}
		}
		queueNode = queueNode[size:]
		queueVal = queueVal[size:]
	}
	return false
}
