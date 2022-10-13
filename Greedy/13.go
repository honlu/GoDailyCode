/*
13、用最少数量的箭引爆气球
2022-10-13
link: 452-https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/
question:
	一支弓箭可以沿着 x 轴从不同点完全垂直地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend，
	且满足  xstart ≤ x ≤ xend，则该气球会被引爆。可以射出的弓箭的数量没有限制。 弓箭一旦被射出之后，可以无限地前进。
	我们想找到使得所有气球全部被引爆，所需的弓箭的最小数量。
	给你一个数组 points ，其中 points [i] = [xstart,xend] ，返回引爆所有气球所必须射出的最小弓箭数。
answer:
	这一题有点类合并序列（数组）的题目。
	首先排序，按照起始节点排序。
	然后依次比较当前节点起始和上一个节点的末端。并且还要更新重叠气球的末端
*/
func findMinArrowShots(points [][]int) int {
	// 排序
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		}
		return false
	})
	if len(points) == 0 {
		return 0
	}
	// 遍历比较
	res := 1                           // points不为空是至少需要一只箭
	for i := 1; i < len(points); i++ { // 里面比较是关键和技巧
		if points[i][0] > points[i-1][1] { // 气球i和i-1不在范围内
			res++
		} else {
			if points[i-1][1] < points[i][1] {
				points[i][1] = points[i-1][1] // 更新到两者都可以被射中的最大范围
			}
		}
	}
	return res
}