package hot100

/*
42-将有序数组转换为二叉搜索
https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/?envType=study-plan-v2&envId=top-100-liked

二叉搜索树的特点：左孩子比根节点小，右孩子比根节点大。
递归将数组的中间节点当做根节点。
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	count := len(nums)
	if count == 0 {
		return nil
	}
	if count == 1 {
		return &TreeNode{Val: nums[0]}
	}
	middle := count / 2
	res := &TreeNode{Val: nums[middle]}
	res.Left = sortedArrayToBST(nums[:middle])
	res.Right = sortedArrayToBST(nums[middle+1:])
	return res
}
