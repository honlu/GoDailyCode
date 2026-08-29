package lcr187

// 以下：模拟，递归。结果：超时，需要优化。
// func iceBreakingGame(num int, target int) int {
// 	temp := make([]int, num)
// 	for i := 0; i < num; i++ {
// 		temp[i] = i
// 	}
// 	var cur func([]int, int) int
// 	cur = func(arr []int, target int) int {
// 		count := len(arr)
// 		if count == 1 {
// 			return arr[0]
// 		}
// 		index := (target - 1) % count
// 		arr = append(arr[index+1:], arr[:index]...)
// 		return cur(arr, target)
// 	}
// 	return cur(temp, target)
// }

// 约瑟夫环递推公式
func iceBreakingGame(n int, m int) int {
	/*
		约瑟夫环问题。
		核心思想：
		每淘汰一个人之后，剩余的人重新编号，
		就变成了一个规模更小、规则完全相同的问题。
		定义：
			f(n) = n 个人进行游戏时，最终留下的人的编号
		当只有 1 个人时：
			f(1) = 0
		假设已经知道 n-1 个人时最终获胜者的位置 f(n-1)，
		那么恢复到 n 个人原来的编号时，需要整体向后偏移 m 个位置：
			f(n) = (f(n-1) + m) % n
		因此可以从：
			1 个人
		→ 2 个人
		→ 3 个人
		→ ...
		→ n 个人
		逐步反推出最终答案。
	*/
	// 只有 1 个人时，唯一的编号是 0，
	// 所以最终留下的人一定是 0。
	idx := 0
	// 从 2 个人开始，逐步计算到 n 个人。
	for size := 2; size <= n; size++ {
		/*
			当前：
				idx = f(size-1)
			现在要求：
				f(size)
			为什么需要 +m？
			因为淘汰一个人之后，会从被淘汰者的下一个位置
			重新开始计数，相当于剩余的人重新进行了编号。
			例如 n=5，m=3：
				原编号：
					0  1  2  3  4
					      ×
					     淘汰
				淘汰 2 后，从 3 开始重新编号：
				原编号：3  4  0  1
				新编号：0  1  2  3
			如果我们已经知道“新编号”中的最终获胜者，
			要恢复成原来的编号，就需要整体向后移动 m 个位置。
			由于所有人围成一个环，所以超过 size 后需要回到开头，
			因此需要对 size 取模：
				f(size) = (f(size-1) + m) % size
		*/
		idx = (idx + m) % size
	}

	return idx
}
