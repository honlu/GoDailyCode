package hot100

/*
43-验证二叉搜索树
https://leetcode.cn/problems/validate-binary-search-tree/description/?envType=study-plan-v2&envId=top-100-liked

中序遍历二叉搜索树，结果是递增的
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var res []int
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		preorder(node.Left)
		res = append(res, node.Val)
		preorder(node.Right)
	}
	preorder(root)
	for i := 0; i < len(res)-1; i++ {
		if res[i] >= res[i+1] { // 注意：不能等于
			return false
		}
	}
	return true
}
