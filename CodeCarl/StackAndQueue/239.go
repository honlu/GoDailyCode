/*
239. 滑动串口最大值
2023-1-13
link: https://leetcode.cn/problems/sliding-window-maximum/
question:
	给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
	你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

	返回 滑动窗口中的最大值 。

answer:
	单调队列的经典题目！怎么写呢？

	难点是如何求一个区间里的最大值呢？ （这好像是废话），暴力一下不就得了。
		暴力方法，遍历一遍的过程中每次从窗口中再找到最大的数值，这样很明显是O(n × k)的算法。
		有的同学可能会想用一个大顶堆（优先级队列）来存放这个窗口里的k个数字，
			这样就可以知道最大的最大值是多少了， 但是问题是这个窗口是移动的，
			而大顶堆每次只能弹出最大值，我们无法移除其他数值，
			这样就造成大顶堆维护的不是滑动窗口里面的数值了。所以不能用大顶堆。

	Golang的暴力方法，会部分示例超时！
	按照carl哥思路：单调队列实现
	队列:放进去窗口里的元素，然后随着窗口的移动，队列也一进一出，每次移动之后，队列告诉我们里面的最大值是什么
	每次窗口移动的时候，调用que.pop(滑动窗口中移除元素的数值)，que.push(滑动窗口添加元素的数值)，然后que.front()就返回我们要的最大值。
	疑惑：队列中的元素一定是排序的，而且最大值再出队口，要不然怎么知道最大值呢。
	如果把窗口里的元素都放进队列里，窗口移动的时候，队列需要弹出元素。
	那么已经排序之后的队列 怎么能把窗口要移除的元素（这个元素可不一定是最大值）弹出呢?
	答：队列没有必要维护窗口里的所有元素，只需要维护有可能成为窗口里最大值的元素就可以了，
		同时保证队列里的元素数值是由大到小的。
		这个维护元素单调递减的队列就叫做单调队列，即单调递减或单调递增的队列。
		C++&Go中没有直接支持单调队列，需要我们自己来实现一个单调队列。
	注意：单调队列不是单纯的给队列中元素排序，那和优先级队列没有什么区别了。

	本体单调队列是递减的！
	设计单调队列的时候，pop，和push操作要保持如下规则：
	pop(value)：如果窗口移除的元素value等于单调队列的出口元素，
		那么队列弹出元素，否则不用任何操作
	push(value)：如果push的元素value大于入口元素的数值，
		那么就将队列入口的元素弹出，直到push元素的数值小于等于队列入口元素的数值为止
*/
// 封装单调队列的结构体
type MyQueue struct {
	queue []int
}

// 初始化
func Constructor() *MyQueue {
	return &MyQueue{
		queue: make([]int, 0),
	}
}

// 获取队头元素
func (this *MyQueue) Front() int {
	return this.queue[0]
}

// 获取队尾元素
func (this *MyQueue) Back() int {
	return this.queue[len(this.queue)-1]
}

// 入队
func (this *MyQueue) Push(val int) {
	// 入队元素如果比队尾元素大，则队尾元素全部删除
	for !this.Empty() && val > this.Back() {
		this.queue = this.queue[:len(this.queue)-1] // 删除队尾元素
	}
	// 当队尾元素大于等于val时，再入队
	this.queue = append(this.queue, val)
}

// 出队: 指定元素的对头元素删除。 关键点：不是传统的队头都出，并且返回元素值
func (this *MyQueue) Pop(val int) {
	if !this.Empty() && val == this.Front() {
		this.queue = this.queue[1:]
	}
}

// 是否为空
func (this *MyQueue) Empty() bool {
	return len(this.queue) == 0
}

// 关键
func maxSlidingWindow(nums []int, k int) []int {
	// 初始化一个单调队列
	queue := Constructor()
	res := make([]int, 0) // 存储每k个最大值的切片结果
	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	res = append(res, queue.Front()) // 记录前k个元素的最大值

	for i := k; i < len(nums); i++ {
		// 滑动窗口往前移动，则要删除i-k位的元素.单调队列中如果存在这个元素要删除
		queue.Pop(nums[i-k]) // !!! 关键要理解
		// 队列插入新的滑动窗口元素
		queue.Push(nums[i])
		// 记录新的最大值
		res = append(res, queue.Front())
	}
	return res
}