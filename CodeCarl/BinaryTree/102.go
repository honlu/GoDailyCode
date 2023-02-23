package binarytree

import "container/list"

/*
5
102、二叉树的层序遍历
day:2022-6-5
Update: 2023-1-15
link: https://leetcode.cn/problems/binary-tree-level-order-traversal/
question:
	给你二叉树的根节点 root ，返回其节点值的 层序遍历 。
	（即逐层地，从左到右访问所有节点）。
idea:
	使用队列实现二叉树广度优先遍历
	借用一个辅助数据结构即队列来实现，队列先进先出，符合一层一层遍历的逻辑，而是用栈先进后出适合模拟深度优先遍历也就是递归的逻辑。
	而这种层序遍历方式就是图论中的广度优先遍历，只不过我们应用在二叉树上。

队列：先进先出。
Go中队列的模拟实现：
	1. 切片实现
	2. container/list
*/
// list模拟队列
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}    // 结果切片
	queue := list.New() // 队列
	queue.PushBack(root)
	var tempArr []int // 层切片
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
			tempArr = append(tempArr, node.Val)
		}
		res = append(res, tempArr) // 加入结果集
		tempArr = []int{}          // 清楚层的数据
	}
	return res
}

// 切片模拟队列
func levelOrder(root *TreeNode) [][]int {
	// 广度有限搜索
	if root == nil {
		return nil
	}
	// 迭代实现
	var res [][]int
	q := []*TreeNode{root} // 一层节点地队列
	for len(q) > 0 {
		temp := []int{} // 一层节点地值
		size := len(q)
		for j := 0; j < size; j++ {
			node := q[j]
			temp = append(temp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, temp) // 添加一层节点地值
		q = q[size:]            // 转换层次，进行下一次的迭代
	}
	return res
}
