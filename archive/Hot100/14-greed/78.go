package greed

/*
78-跳跃游戏
*/
func canJump(nums []int) bool {
	// 能接触到的索引下标的最大跳跃覆盖范围
	if len(nums) == 1 {
		return true
	}
	index := 0
	maxCoverr := index + nums[index]
	for i := index + 1; i <= maxCoverr; i++ {
		if i+nums[i] > maxCoverr {
			maxCoverr = i + nums[i]
		}
		if maxCoverr >= len(nums)-1 {
			return true
		}
	}
	return false
}

/*
func canJump(nums []int) bool {
    // 递归+记忆性（避免重复）
    memo := make([]int, len(nums)) // 0: 未访问, 1: 可达, -1: 不可达

    var dfs func(pos int) bool
    dfs = func(pos int) bool {
        if pos >= len(nums)-1 {
            return true
        }
        if memo[pos] != 0 {
            return memo[pos] == 1
        }
        maxJump := nums[pos]
        for i := 1; i <= maxJump; i++ {
            if dfs(pos + i) {
                memo[pos] = 1
                return true
            }
        }
        memo[pos] = -1
        return false
    }

    return dfs(0)
}

*/
