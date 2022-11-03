/*
221. 最大正方形-中等
2022-11-3
link: https://leetcode.cn/problems/maximal-square
question:
	在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，
	并返回其面积。
answer:
	动态规划五步曲：
	1、dp数组和下标含义
		dp[i][j]表示在矩阵中以(i,j)为右下角，且只包含1的正方形的边长最大值。
		则其平方就是最大正方形的面积。
	2、递推公式
		两种情况
		如果matrix[i][j]=0,则dp[i][j]=0
		如果matrix[i][j]=1,则dp[i][j]的值由其上方，左方和左上方的三个相邻位置的dp值决定。
		即当前位置的元素值等于三个相邻位置的元素中的最小值加1.
		dp[i][j] = min(dp[i-1][j-1],dp[i-1][j],dp[i][j-1])+1
	3、初始化
		dp[i][j]=int(matrix[i][j])
	4、遍历顺序
		通过递推公式发现，从上到下，从左到
	5、举例推导
*/
func maximalSquare(matrix [][]byte) int {
	// dp数组和初始化
	dp := make([][]int, len(matrix))
	res := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))   // 注意这里长度
		for j := 0; j < len(matrix[i]); j++ { // 注意这里长度
			dp[i][j] = int(matrix[i][j] - '0') // 初始化
			if dp[i][j] == 1 {
				res = 1
			}
		}
	}
	// 遍历顺序
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1 // 递推公式！
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res * res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
func maximalSquare(matrix [][]byte) int {
	// dp数组和初始化
	dp := make([][]int, len(matrix))
	res := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))  // 注意这里长度
		dp[i][0] = int(matrix[x][y]-'0')
		for j := 0; j < len(matrix[i]); j++ { // 注意这里长度
			if matrix[i][j] == '1' {
				res = 1
			}
		}
	}
	for j := 0; j < len(matrix[i]); j++ { // 注意这里长度
		if matrix[i][j] == '1' {
			res = 1
		}
	}
	// 遍历顺序
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1 // 递推公式！
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res * res
}


*/