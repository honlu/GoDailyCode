/*
6、不同路径2
2022-10-14
link: 63-https://leetcode.cn/problems/unique-paths-ii/
question:
	一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
	机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
	现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
answer:
	相比于“不同路径”问题增加了一个条件，网格中有障碍物！
	如何处理障碍物呢？就是标记对应的dp table（dp数组）保持初始值(0)就可以了。
		或者说没有障碍物时，再更新dp[i][j]
	1、dp[i][j]的定义：表示从（0 ，0）出发到达第（i，j）个位置有多少条路径。
	2、递推公式：dp[i][j] = dp[i - 1][j]+dp[i][j-1];因为只有这两个方向过来。
	3、dp的初始化：首先dp[i][0]一定都是1，因为从(0, 0)的位置到(i, 0)的路径只有一条，
		注意若存在一个为障碍物，则这一行后面的初始值要为0.那么dp[0][j]也同理
	4、确定遍历顺序：这里要看一下递归公式dp[i][j] = dp[i - 1][j] + dp[i][j - 1]，dp[i][j]都是从其上方和左方推导而来，那么从左到右一层一层遍历就可以了。
	5、草稿纸：举例推导dp是否正确
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始化
	for i := 0; i < n; i++ { // dp[0][i]
		if obstacleGrid[0][i] == 1 { // 有障碍物时，后面都要为0
			break
		}
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ { // dp[i][0]
		if obstacleGrid[i][0] == 1 { // 有障碍物时，后面都要为0
			break
		}
		dp[i][0] = 1
	}
	// 遍历，更新dp矩阵
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 { // 没有障碍物，再更新dp[i][j];有障碍物，就保持dp[i][j]的初始值0
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}

		}
	}
	return dp[m-1][n-1]
}