/*
14. 重新安排行程-困难
day：2022-10-8
link:332-https://leetcode.cn/problems/reconstruct-itinerary/
question:
	给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，
	对该行程进行重新规划排序。所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，
	所以该行程必须从 JFK 开始。
answer:
	更像是图论中的深度优先搜索。
	实际上确实是深搜，但这是深搜中使用了回溯的例子，在查找路径的时候，
	如果不回溯，怎么能查到目标路径呢。

以下改变官方，非代码随想录思路。
Undo: 示例可以通过，但是超出内存了？？？
*/
func findItinerary(tickets [][]string) []string {
	var res []string
	m := map[string][]string{} // 每个起点对应的终点数组
	// 初始化，将图的邻接表存进字典
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		m[src] = append(m[src], dst)
	}
	// 按字典序排序
	for key := range m {
		sort.Strings(m[key]) // 把每个起点对应的数组按照字母排序
	}
	backTrack("JFK", m, &res)
	// 翻转数组——这一步不能少！
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res

}

func backTrack(cur string, m map[string][]string, res *[]string) {
	for {
		// 这个起点终点已经走完了，或者是个死胡同没有回去的路
		if v, ok := m[cur]; !ok || len(v) == 0 {
			break
		}
		// 还有路，继续走
		temp := m[cur][0]
		backTrack(temp, m, res)
		// 去掉这条路
		m[cur] = m[cur][1:] // 注意函数内，还可以改变吗
	}
	// 往下递归到终点之后开始入栈
	// 往上递归返回的入栈就是来时的路反着走回去，这样得到的数组是 路径的倒序
	// 越在链路后面的先入栈
	*res = append(*res, cur)
}