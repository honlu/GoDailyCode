package main

// 二叉搜索树转为双向链表
// Python3 版本:
/*
class Solution:
    def treeToDoublyList(self, root: "Node") -> "Node":
        if not root:
            return None
        pre = None
        head = None
        def dfs(node):
            nonlocal pre, head
            if node is None:
                return
            # 中序遍历
            dfs(node.left)
            if pre:
                pre.right = node
            else:
                head = node
            node.left = pre
            pre = node
            dfs(node.right)

        dfs(root)
        # 变成循环链表（可选部分），即首尾相连
        head.left = pre
        pre.right = head
        return head
*/
func treeToDoublyList(root *TreeNode) *TreeNode {
	var pre, head *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 中序遍历
		dfs(node.Left)
		if pre != nil {
			pre.Right = node
		} else {
			head = node
		}
		node.Left = pre
		pre = node
		dfs(node.Right)
	}
	dfs(root)
	return head
}
