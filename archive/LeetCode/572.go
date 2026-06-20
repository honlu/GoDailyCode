/*
572. 另一棵树的子树
2023-3-11
link: https://leetcode.cn/problems/subtree-of-another-tree/
question:
	给你两棵二叉树 root 和 subRoot 。检验 root 中是否包含和 subRoot 具有相同结构和节点值的子树。
	如果存在，返回 true ；否则，返回 false 。
	二叉树 tree 的一棵子树包括 tree 的某个节点和这个节点的所有后代节点。tree 也可以看做它自身的一棵子树。
answer:
	可以递归和迭代实现。这里使用队列，迭代实现
	即首先找到root树中和subRoot根节点一样的值，然后判断两个子树是否相同
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// 树的遍历
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			temp := queue[i]
			if temp.Val == subRoot.Val && isSame(temp, subRoot) {
				return true
			}
			if temp.Left != nil {
				queue = append(queue, temp.Left)
			}
			if temp.Right != nil {
				queue = append(queue, temp.Right)
			}
		}
		queue = queue[size:]
	}
	return false
}

func isSame(root *TreeNode, subRoot *TreeNode) bool {
	queue1 := []*TreeNode{root}
	queue2 := []*TreeNode{subRoot}
	for len(queue1) > 0 && len(queue2) > 0 {
		size := len(queue1)
		for i := 0; i < size; i++ {
			// 下面处理逻辑比较重要
			if queue1[i] == nil && queue2[i] == nil {
				continue
			} else if queue1[i] == nil || queue2[i] == nil || queue1[i].Val != queue2[i].Val {
				return false
			} else {
				queue1 = append(queue1, queue1[i].Left, queue1[i].Right)
				queue2 = append(queue2, queue2[i].Left, queue2[i].Right)
			}
		}
		queue1 = queue1[size:]
		queue2 = queue2[size:]
	}
	return true
}

// 迭代实现代码比较多。这里递归实现，以供参考。
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	//对根节点进行判断是否为true
	if isSameTree(s, t) { //这是一个很苛刻的条件，不仅根节点相同，左子树，右子树均相同
		return true
	}
	//对左右子树判断其中一个为true就为true
	return isSubtree(s.Left, t) || isSubtree(s.Right, t) //有一个为true就为true，这个更是包含了上面那个条件，更更苛刻

}
func isSameTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	return s.Val == t.Val && isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}
