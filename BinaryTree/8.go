package binarytree

import "container/list"

/*
8、二叉树的最大深度
day :2022-6-5
*/
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 递归实现树的最大高度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// 迭代实现-层次遍历
func maxDepthIter(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()         // 保存当前层的长度，然后处理当前层。不可再for里面使用queue.Len()，防止添加元素动态变化
		for i := 0; i < size; i++ { // 依次遍历当前层在队列的头部元素
			node := queue.Remove(queue.Front()).(*TreeNode) // 头部出队列，删除元素，返回元素的值，转为TreeNode指针
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res += 1
	}
	return res
}
