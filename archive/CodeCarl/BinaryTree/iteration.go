package binarytree

import "container/list"

/*
3、二叉树的迭代遍历
day:2022-6-5
idea:递归的实现就是：每一次递归调用都会把函数的局部变量、参数值和返回地址等压入调用栈中，
然后递归返回的时候，从栈顶弹出上一次递归的各项参数，
所以这就是递归为什么可以返回上一层位置的原因。

这里栈使用列表来实现，只是使用栈的思想：元素先进后出
*/

// 前序遍历-迭代
func preorderIterateTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}
	stack := list.New()  // 初始化列表，list内部的实现原理是双链表
	stack.PushBack(root) // 后方插入元素

	for stack.Len() > 0 { // 每次都要重新判断栈中元素是否为空
		node := stack.Remove(stack.Back()).(*TreeNode) // Back()获取列表尾结点，Remove()删除元素并且返回被删除的元素的值
		res = append(res, node.Val)                    // 中节点先处理
		if node.Right != nil {
			stack.PushBack(node.Right) // 右孩子先入栈，后处理
		}
		if node.Left != nil {
			stack.PushBack(node.Left) // 左孩子后入栈，比右孩子先处理
		}
	}
	return res
}

// 中序遍历-迭代
//用迭代法写出了二叉树的前后中序遍历，大家可以看出前序和中序是完全两种代码风格，并不像递归写法那样代码稍做调整，就可以实现前后中序。
//这是因为前序遍历中访问节点（遍历节点）和处理节点（将元素放进result数组中）可以同步处理，但是中序就无法做到同步！
func inorderIterateTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}
	stack := list.New()
	node := root

	for node != nil || stack.Len() > 0 { // 当前结点不为空或者栈不为空
		if node != nil { // 先迭代访问最底层的左子树结点
			stack.PushBack(node)
			node = node.Left
		} else { // 到达最左结点后，处理栈顶结点
			// 注意这里 = 和 := 结果完全不一样？为什么呢？
			node = stack.Remove(stack.Back()).(*TreeNode) // Back()获取列表尾结点，Remove()删除元素并且返回被删除的元素的值
			res = append(res, node.Val)
			node = node.Right // 取栈顶元素右结点
		}
	}
	return res
}

// 后序遍历-迭代
func postorderIteratTraversal(root *TreeNode) (res []int) {
	if root == nil {
		return res
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, node.Val) // 中结点先处理
		if node.Left != nil {       // 左孩子先入栈
			stack.PushBack(node.Left)
		}
		if node.Right != nil { // 右孩子后入栈，会先被处理
			stack.PushBack(node.Right)
		}
	}
	// 上面得到的是中-右-左，后序遍历是左-右-中，所以需要将最终的数组反转
	return reverse(res)
}

func reverse(res []int) []int {
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l, r = l+1, r+1
	}
	return res
}
