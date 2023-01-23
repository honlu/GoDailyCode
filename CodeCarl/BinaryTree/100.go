package binarytree

/*
13
100、相同树
day:2022-6-6
link: https://leetcode.cn/problems/same-tree/
question:
	给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
	如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
idea:
	认清判断对称树本质之后，对称树的代码稍作修改就可以直接用来AC
	也是相对应节点匹配对比！

	递归版本可以，两个节点对比实现。


	迭代版本，利用层次遍历，左边树层进，右边树和层对比，相同进入下一层对比，否则不相同.
	使用两个队列分别存储两个二叉树的节点。初始时将两个二叉树的根节点分别加入两个队列。每次从两个队列各取出一个节点，进行如下比较操作。
		比较两个节点的值，如果两个节点的值不相同则两个二叉树一定不同；
		如果两个节点的值相同，则判断两个节点的子节点是否为空，如果只有一个节点的左子节点为空，
			或者只有一个节点的右子节点为空，则两个二叉树的结构不同，因此两个二叉树一定不同；
		如果两个节点的子节点的结构相同，
			则将两个节点的非空子节点分别加入两个队列，子节点加入队列时需要注意顺序，如果左右子节点都不为空，则先加入左子节点，后加入右子节点。
	如果搜索结束时两个队列同时为空，
			则两个二叉树相同。如果只有一个队列为空，则两个二叉树的结构不同，因此两个二叉树不同。


测试用例：
[1,2,3]
[1,2,3]
[1,2]
[1,null,2]
[1,2,1]
[1,1,2]
[1]
[1,null,2]
*/

// 递归实现(深度优先搜索)
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 递归实现
	switch {
	case p == nil && q == nil:
		return true
	case p == nil || q == nil:
		return false
	case p.Val != q.Val:
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 迭代版本(层次遍历,两个队列)
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	// 两个队列(层次遍历)
	if p.Val != q.Val {
		{
			return false
		}
	}
	queue1, queue2 := []*TreeNode{p}, []*TreeNode{q}
	for len(queue1) > 0 && len(queue2) > 0 {
		size := len(queue1)
		for i := 0; i < size; i++ {
			// 开始添加新的一层节点
			if queue1[i].Left != nil && queue2[i].Left != nil && queue1[i].Left.Val == queue2[i].Left.Val {
				queue1 = append(queue1, queue1[i].Left)
				queue2 = append(queue2, queue2[i].Left)
			} else if queue1[i].Left == nil && queue2[i].Left == nil {
				// 这里不执行代码，只是为了else方便
			} else {
				return false
			}

			if queue1[i].Right != nil && queue2[i].Right != nil && queue1[i].Right.Val == queue2[i].Right.Val {
				queue1 = append(queue1, queue1[i].Right)
				queue2 = append(queue2, queue2[i].Right)
			} else if queue1[i].Right == nil && queue2[i].Right == nil {
				// 这里不执行代码
			} else {
				return false
			}
		}
		queue1 = queue1[size:]
		queue2 = queue2[size:]
	}
	// 最后还有特殊情况，比如一棵树结束，另外一个树还没结束
	// [1]
	// [1,null,2]
	return len(queue1) == 0 && len(queue2) == 0
}
