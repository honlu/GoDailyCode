package binarytree

import (
	"container/list"
	"strconv"
)

/*
12、二叉树的所有路径
day:2022-6-6
idea:
要求从根节点到叶子的路径，所以需要前序遍历，这样才方便让父节点指向孩子节点，
找到对应的路径。
在这道题目中将第一次涉及到回溯，因为我们要把路径记录下来，
需要回溯来回退一一个路径在进入另一个路径。
先使用递归的方式，来做前序遍历。要知道递归和回溯就是一家的，本题也需要回溯。

回溯和递归是一一对应的，有一个递归，就要有一个回溯，这么写的话相当于把递归和回溯拆开了， 一个在花括号里，一个在花括号外。
所以回溯要和递归永远在一起，世界上最遥远的距离是你在花括号里，而我在花括号外！

本题是第一次涉及回溯！回溯和递归是相伴相生！
对于本题充分了解递归与回溯的过程之后，有精力的同学可以在去实现迭代法。
*/

// 递归+隐形回溯-前序遍历
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	var traversal func(node *TreeNode, s string)
	traversal = func(node *TreeNode, s string) {
		if node.Left == nil && node.Right == nil { // 遇到叶子节点
			temp := s + strconv.Itoa(node.Val) // 记录最后一个节点（叶子节点）
			res = append(res, temp)            // 收集一个路径
			return
		}
		s = s + strconv.Itoa(node.Val) + "->" // 前序遍历，先处理中间结点，将其放进路径中
		if node.Left != nil {
			traversal(node.Left, s) // 递归。注意这里s是复制赋值，不是引用。所以执行完递归函数之后，s依然是之前的数值（相当于回溯了）
		}
		if node.Right != nil {
			traversal(node.Right, s) // 递归
		}
	}
	traversal(root, "")
	return res
}

// 迭代实现-前序遍历
func binaryTreePathsIter(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	stack := list.New() // 初始化栈，保存树的遍历节点
	path := list.New()  // 保存遍历路径的节点
	res := []string{}   // 保存最终路径集合
	stack.PushBack(root)
	path.PushBack(strconv.Itoa(root.Val))

	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode) // 取出节点：中。出栈、删除、返回删除的Value接口转为节点
		temp := path.Remove(path.Back()).(string)      // 取出该节点对应的路径
		if node.Left == nil && node.Right == nil {     // 遇到叶子节点
			res = append(res, temp)
		}

		if node.Right != nil {
			stack.PushBack(node.Right) // 右孩子先入栈，后处理
			path.PushBack(temp + "->" + strconv.Itoa(node.Right.Val))
		}
		if node.Left != nil {
			stack.PushBack(node.Left) // // 左孩子后入栈，比右孩子先处理.利用栈的先进后出思想
			path.PushBack(temp + "->" + strconv.Itoa(node.Left.Val))
		}
	}
	return res
}
