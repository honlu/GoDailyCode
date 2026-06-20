/*
384. 打乱数组
2023-3-13
link: https://leetcode.cn/problems/shuffle-an-array/
question:
	给你一个整数数组 nums ，设计算法来打乱一个没有重复元素的数组。打乱后，数组的所有排列应该是 等可能 的。
	实现 Solution class:
		Solution(int[] nums) 使用整数数组 nums 初始化对象
		int[] reset() 重设数组到它的初始状态并返回
		int[] shuffle() 返回数组随机打乱后的结果
answer:
	打乱后和打乱前肯定要用两个数组来保存的！
	知识这里打乱使用的是暴力随机打乱，每次生成一个随机索引，将原来数组对应值赋值给新的元素，并且删除这个索引。
	然后继续相同操作
*/
type Solution struct {
	nums    []int // 当前数组
	origins []int // 原先数组
}

func Constructor(nums []int) Solution {
	return Solution{nums, append([]int{}, nums...)}
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.origins) // 将origins赋值到原来的数组
	return this.nums
}

func (this *Solution) Shuffle() []int {
	n := len(this.nums)
	temp := make([]int, n)
	// 随机选一个nums里元素赋值给打乱后的数组
	for i := 0; i < n; i++ {
		j := rand.Intn(len(this.nums))
		temp[i] = this.nums[j]
		this.nums = append(this.nums[:j], this.nums[j+1:]...) // 注意j+1
	}
	this.nums = temp
	return this.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */