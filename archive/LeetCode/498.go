/*
498. 对角线遍历
2022-12-25
link: https://leetcode.cn/problems/diagonal-traverse/
question:
	给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。
answer:
https://leetcode.cn/problems/diagonal-traverse/solution/dui-jiao-xian-bian-li-by-leetcode-soluti-plz7/
	直接模拟遍历。首先难点要发现规律并总结：
	1. 一共m+n-1条对角线。相邻对角线方向不一样，先从左下到右上，然后右上到左下，再次循环。
	2. 设对角线从上到下的编号是i：[0,m-n-2]。当i为偶数时，则第i条对角线的走向时从下到上遍历，当i为奇数时，则返过来。
	3. 当第i条对角线从下往上遍历时，每次行索引减1，列索引加1，直到矩阵的边缘为止：
		- 当i<m时，对角线起点位置为(i,0)
		- 当i>=m时，对角线起点位置为(m-1,i-m+1)
	4. 当第i条对角线从上到下遍历时，每次行索引加1，列索引减1，直到矩阵的边缘为止：
		- 当i<n时，对角线起点位置为(0,i)
		- 当i>=n时，对角线起点位置为(i-n+1,n-1)
*/
func findDiagonalOrder(mat [][]int) []int {
	var res []int
	var m, n int = len(mat), len(mat[0])
	for i := 0; i < m+n-1; i++ {
		var startI, startJ int
		if i%2 == 0 { // 从上到下遍历
			if i < m {
				startI, startJ = i, 0
			} else {
				startI, startJ = m-1, i-m+1
			}
			for startI >= 0 && startJ <= n-1 {
				res = append(res, mat[startI][startJ])
				startI, startJ = startI-1, startJ+1
			}
		} else { // 从下到上遍历
			if i < n {
				startI, startJ = 0, i
			} else {
				startI, startJ = i-n+1, n-1
			}
			for startI <= m-1 && startJ >= 0 {
				res = append(res, mat[startI][startJ])
				startI, startJ = startI+1, startJ-1
			}
		}
	}
	return res
}