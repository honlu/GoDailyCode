package binarytree

import "container/list"

/*
6、翻转二叉树
day:2022-6-5
update: 2023-1-15
link:LeetCode-226 https://leetcode.cn/problems/invert-binary-tree/
question:
	给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
idea:
注意只要把每一个节点的左右孩子翻转一下，就可以达到整体翻转的效果
这道题目使用前序遍历和后序遍历都可以，唯独递归-中序遍历不方便，因为中序遍历会把某些节点的左右孩子翻转了两次！建议拿纸画一画，就理解了
那么层序遍历可以不可以呢？依然可以的！只要把每一个节点的左右孩子翻转一下的遍历方式都是可以的！
*/

// 翻转-递归版本的前序遍历 中左右
func inverTreePre(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	inverTreePre(root.Left)  // 递归-左
	inverTreePre(root.Right) // 递归-右
	return root
}

// 翻转-递归版本的后序遍历  左右中
func inverTreePost(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	inverTreePost(root.Left)                      // 遍历左结点
	inverTreePost(root.Right)                     // 遍历右结点
	root.Left, root.Right = root.Right, root.Left // 交换
	return root
}

// 翻转-迭代版本的前序遍历，栈实现
func inverTreePreIter(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		node.Left, node.Right = node.Right, node.Left
		if node.Right != nil {
			stack.PushBack(node.Right) // 右孩子先入栈，后处理
		}
		if node.Left != nil {
			stack.PushBack(node.Left) // 左孩子后入栈，比右孩子先处理
		}
	}
	return root
}

// 翻转-迭代统一版本的后序遍历，栈实现
func inverTreePostIter(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		node.Left, node.Right = node.Right, node.Left // 中结点先处理
		if node.Left != nil {                         // 左孩子先入栈
			stack.PushBack(node.Left)
		}
		if node.Right != nil { // 右孩子后入栈，会先被处理
			stack.PushBack(node.Right)
		}
	}
	return root
}

// 翻转-迭代统一版本的中序遍历，栈实现
// 以下是错误代码，可能某些层会翻转两次，这是为什么呢？
// func inverTreeInIter(root *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}
// 	stack := list.New()
// 	node := root

// 	for node != nil || stack.Len() > 0 { // 当前结点不为空或者栈不为空
// 		if node != nil { // 先迭代访问最底层的左子树结点
// 			stack.PushBack(node)
// 			node = node.Left
// 		} else { // 到达最左结点后，处理栈顶结点
// 			node = stack.Remove(stack.Back()).(*TreeNode) // Back()获取列表尾结点，Remove()删除元素并且返回被删除的元素的值
// 			node.Left, node.Right = node.Right, node.Left // 交换
// 			node = node.Right                             // 取栈顶元素右结点
// 		}
// 	}
// 	return root
// }

// 翻转-迭代统一版本的层序遍历，队列实现
func inverTreeLeverl(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := list.New() // 队列
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()         // 保存当前层的长度，然后处理当前层。不可再for里面使用queue.Len()，防止添加元素动态变化
		for i := 0; i < size; i++ { // 依次遍历当前层在队列的头部元素
			node := queue.Remove(queue.Front()).(*TreeNode) // 头部出队列，删除元素，返回元素的值，转为TreeNode指针
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}

		}
	}
	return root
}
