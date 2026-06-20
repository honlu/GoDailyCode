package dynamicprogramming

/*
82-118：杨辉三角
https://leetcode.cn/problems/pascals-triangle/description/?envType=study-plan-v2&envId=top-100-liked
*/
func generate(numRows int) [][]int {
	// res[i][j] = res[i-1][j-1] + res[i-1][j]
	var res [][]int
	for i := 0; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0], temp[i] = 1, 1
		res = append(res, temp)
		for j := 1; i >= 2 && j < i; j++ { // 注意中间判断条件
			res[i][j] = res[i-1][j-1] + res[i-1][j] // 公式
		}

	}
	return res
}
