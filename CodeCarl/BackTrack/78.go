/*
9
78. 子集
day: 2022-10-8
link: https://leetcode.cn/problems/subsets/
question:
	给定一组不含重复元素的整数数组 nums，
	返回该数组所有可能的子集（幂集）。
answer:
	注意：组合、分割和子集的不同。
	如果把 子集问题、组合问题、分割问题都抽象为一棵树的话，
	那么组合问题和分割问题都是收集树的叶子节点，
	而子集问题是找树的所有节点！（遍历这个树的时候，把所有节点都记录下来，就是要求的子集集合。）

	子集也是一种组合问题，因为它的集合是无序的，子集{1,2} 和 子集{2,1}是一样的。
	那么既然是无序，取过的元素不会重复取，写回溯算法的时候，
	for就要从startIndex开始，而不是从0开始！
*/
// 回溯
func subsets(nums []int) [][]int {
	var res [][]int // 存结果集
	var path []int  //  为子集收集元素
	sort.Ints(nums) // 避免重复
	// 回溯定义
	var backTrack func(start int)
	backTrack = func(start int) {
		// 每个节点都要添加到结果中
		res = append(res, append([]int{}, path...))

		// 子集好像不需要base case。这个顺序一定要在res添加之后。
		if start == len(nums) {
			return
		}

		// 回溯标准算法步骤
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i]) // 处理
			backTrack(i + 1)             // 递归
			path = path[:len(path)-1]    // 回溯
		}
	}
	backTrack(0)
	return res
}

/* 回溯 */
func subsets(nums []int) [][]int {
	var res [][]int
	var track []int
	sort.Ints(nums)
	backTrack(0, nums, track, &res)
	return res
}

// 递归
func backTrack(startIndex int, nums, track []int, res *[][]int) {
	// 每个节点都添加到结果中
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp) // 注意这里*

	for i := startIndex; i < len(nums); i++ {
		// if i> startIndex && nums[i]==nums[i-1]{  // 重复元素排除，题目已经说明，所以这段代码可以不用要了
		// 	continue // 这里是要continue还是break呢？如果nums有重复元素
		// }
		track = append(track, nums[i])
		backTrack(i+1, nums, track, res)
		// 回溯
		track = track[:len(track)-1]
	}
}