/*
129. 求根节点到叶结点数字之和
2022-12-4
link: https://leetcode.cn/problems/sum-root-to-leaf-numbers/
question:
	给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
	每条从根节点到叶节点的路径都代表一个数字：
	例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
	计算从根节点到叶节点生成的 所有数字之和 。
answer:
	前序遍历变形
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var ans int
	var traversal func(node *TreeNode, temp string)
	traversal = func(root *TreeNode, temp string) {
		if root == nil { // 到末尾结束
			return
		}
		// 前序遍历，中间节点处理
		temp += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil { // 当到达叶子节点时，进行加法
			// fmt.Println(temp)
			num, _ := strconv.Atoi(temp)
			ans += num
			// return // 可忽略
		}
		traversal(root.Left, temp)  // 递归左节点
		traversal(root.Right, temp) // 递归右节点
		// 回溯
		temp = temp[:len(temp)-1]
	}
	traversal(root, "")
	return ans
}
