/*
9. 子集
day: 2022-10-8
link: 78-https://leetcode.cn/problems/subsets/
question:
	给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
answer:
	如果把 子集问题、组合问题、分割问题都抽象为一棵树的话，
	那么组合问题和分割问题都是收集树的叶子节点，
	而子集问题是找树的所有节点！
*/
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
		// 	continue
		// }
		track = append(track, nums[i])
		backTrack(i+1, nums, track, res)
		// 回溯
		track = track[:len(track)-1]
	}
}