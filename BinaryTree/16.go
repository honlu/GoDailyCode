package binarytree

/*
18、路径总和
day:2022-6-7
link:https://leetcode.cn/problems/path-sum/
link2:https://leetcode.cn/problems/path-sum-ii/
idea:
路径总和：要遍历从根节点到叶子节点的的路径看看总和是不是目标和。
路径总和ii：要遍历整个树，找到所有路径，所以递归函数不要返回值！
*/

// 递归实现路经总和：1还比较简单，只是判断是否存在
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val                                        // 将targetSum在遍历每层的时候都减去本层节点的值
	if root.Left == nil && root.Right == nil && targetSum == 0 { // 如果剩余的targetSum为0，则正好是符合的结果
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum) // 否则递归继续找。
}

// 递归实现：返回所有存在的路径.
// 深度优先搜索（前中序遍历都可以，这里采用前序）：枚举每一条从根节点到叶子节点的路径。当我们遍历到叶子节点，且此时路径和恰为目标和时，我们就找到了一条满足条件的路径。
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
