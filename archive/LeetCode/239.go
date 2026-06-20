/*
239. 滑动窗口最大值
2022-11-27
link: https://leetcode.cn/problems/sliding-window-maximum/
question:
	给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
	你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
	返回 滑动窗口中的最大值 。
answer:
	模拟操作遍历！但注意超时问题，即测试数据过长！
	对于每个滑动窗口，我们可以使用 O(k)的时间遍历其中的每一个元素，找出其中的最大值。
	对于长度为 n 的数组 nums 而言，窗口的数量为 n-k+1，
	因此该算法的时间复杂度为 O((n−k+1)k)=O(nk)，会超出时间限制，因此我们需要进行一些优化。
	我们可以想到，对于两个相邻（只差了一个位置）的滑动窗口，
	它们共用着 k-1 个元素，而只有 1 个元素是变化的。我们可以根据这个特点进行优化。
	思路参考下面：
https://leetcode.cn/problems/sliding-window-maximum/solution/hua-dong-chuang-kou-zui-da-zhi-by-leetco-ki6m/
*/

// 优先队列（堆），自己暂时写不出来！
var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q) // container/heap 这个包来实现堆的操作。堆实际上是一个树的结构，本解法是使用大根堆。

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

// O(n*k)超时版，大多数人都会写。然而改进对自己有点难，加油！
// func maxSlidingWindow(nums []int, k int) []int {
// 	var res []int
// 	if k == 1 {
// 		return nums
// 	}
// 	temp := nums[0,k]
// 	m := max(temp)
// 	res = append(res, m)
// 	for i := 1; i <= len(nums)-k; i++ {
// 		if len(i+k-1) > m {

// 		}
// 		res = append(res, m)
// 	}
// 	return res
// }

// func max(temp []int) int {
// 	m := temp[0]
// 	for i := 1; i < len(temp); i++ {
// 		if temp[i] > m {
// 			m = temp[i]
// 		}
// 	}
// 	return m
// }