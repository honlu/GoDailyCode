package binarytree

/*
18
654、最大二叉树
day:2022-6-9
update:2023-1-25
question:
	给定一个不重复的整数数组 nums 。
	最大二叉树 可以用下面的算法从 nums 递归地构建:

	创建一个根节点，其值为 nums 中的最大值。
	递归地在最大值 左边 的 子数组前缀上 构建左子树。
	递归地在最大值 右边 的 子数组后缀上 构建右子树。
	返回 nums 构建的 最大二叉树 。
link: https://leetcode.cn/problems/maximum-binary-tree/
idea:
	构造树一般采用的是前序遍历，因为先构造中间节点，然后递归构造左子树和右子树。
	注意类似用数组构造二叉树的题目，每次分隔尽量不要定义新的数组，而是通过下标索引直接在原数组上操作，
	这样可以节约时间和空间上的开销。
	一般情况来说：如果让空节点（空指针）进入递归，就不加if，如果不让空节点进入递归，
	就加if限制一下， 终止条件也会相应的调整。

递归的方法比较好理解，还有单调栈的解法，比较难理解。
*/

// 递归，前序遍历方式构造二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil // 这一步挺重要，递归结束的条件
	}
	// 首先从一个数组中找到最大值的索引
	index := findMaxIndex(nums)
	// 然后构造二叉树
	root := &TreeNode{
		Val:   nums[index],                                // 先构建中间节点
		Left:  constructMaximumBinaryTree(nums[:index]),   // 左半边
		Right: constructMaximumBinaryTree(nums[index+1:]), // 右半边
	}
	return root
}

// 从数组中找到最大值并返回其的索引
func findMaxIndex(nums []int) int {
	var index int
	for i, j := 0, len(nums); i < j; i++ {
		if nums[i] > nums[index] {
			index = i
		}
	}
	return index
}
