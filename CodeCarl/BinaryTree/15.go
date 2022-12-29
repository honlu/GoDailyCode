package binarytree

import "container/list"

/*
15、找树左下角的值
day:2022-6-6
link:https://leetcode.cn/problems/find-bottom-left-tree-value/
idea:本地要找出树的最后一行找到最左边的值。
此时大家应该想起用层序遍历是非常简单的了，反而用递归的话会比较难一点。
我们依然还是先介绍递归法。
*/

// 递归法——前序遍历
// 如果使用递归法，如何判断是最后一行呢，其实就是深度最大的叶子节点一定是最后一行。
// 如何找最左边的呢？可以使用前序遍历，这样才先优先左边搜索，然后记录深度最大的叶子节点，此时就是树的最后一行最左边的值。
// 如果需要遍历整棵树，递归函数就不能有返回值。如果需要遍历某一条固定路线，递归函数就一定要有返回值！
// 不能跑通过，问题在哪里呢？ 问题出在递归函数里面node!!千万别写成root了！
func findBottomLeftValue(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	res := 0
	maxDeep := 0
	var findLeftValue func(node *TreeNode, deep int) // deep 当前深度或层度
	findLeftValue = func(node *TreeNode, deep int) {
		if node.Left == nil && node.Right == nil { // 是叶子节点
			if deep > maxDeep { // 是最新一层
				res = root.Val
				maxDeep = deep
			}
		}
		if node.Left != nil {
			// 递归和回溯
			findLeftValue(root.Left, deep+1)
		}
		if node.Right != nil {
			findLeftValue(root.Right, deep+1)
		}
	}
	findLeftValue(root, 0)
	return res
}

// 迭代-层次遍历输出结果
func findBottomLeftValueIter(root *TreeNode) int {
	queue := list.New()
	res := 0
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode) // 弹出队列头部，删除并返回其值
			if i == 0 {                                     // 层的左边第一个元素
				res = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return res
}
