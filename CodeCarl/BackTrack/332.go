/*
14
332. 重新安排行程-困难
day：2022-10-8
update: 2023-2-6 by lu
link:332-https://leetcode.cn/problems/reconstruct-itinerary/
question:
	给定一个机票的字符串二维数组 [from, to]，子数组中的两个成员分别表示飞机出发和降落的机场地点，
	对该行程进行重新规划排序。所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，
	所以该行程必须从 JFK 开始。
answer:
	深搜+回溯
	更像是图论中的深度优先搜索。
	实际上确实是深搜，但这是深搜中使用了回溯的例子，在查找路径的时候，
	如果不回溯，怎么能查到目标路径呢。

	首先先把图的邻接表存进字典，并且按字典序排序，然后从‘JFK’开始深搜，
	每前进一层就减去一条路径，直到某个起点不存在路径的时候就会跳出while循环进行回溯，
	相对先找不到路径的一定是放在相对后面，
	所以当前搜索的起点from会插在当前输出路径的第一个位置。

	Carl这道题目有几个难点：
	- 一个行程中，如果航班处理不好容易变成一个圈，成为死循环
		出发机场和到达机场也会重复的，如果在解题的过程中没有对集合元素处理好，就会死循环。
	- 有多种解法，字母序靠前排在前面，让很多同学望而退步，如何记录映射关系呢 ？
		map存储，但出发机场和到达机场是会重复的，搜索的过程没及时删除目的机场就会死循环。
	- 使用回溯法（也可以说深搜） 的话，那么终止条件是什么呢？
	- 搜索的过程中，如何遍历一个机场所对应的所有机场。

	如果单纯的回溯搜索（深搜）并不难，难还难在容器的选择和使用上。
*/
// carl代码
type pair struct {
	target  string
	visited bool
}
type pairs []*pair

func (p pairs) Len() int {
	return len(p)
}
func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p pairs) Less(i, j int) bool { // 方便排序
	return p[i].target < p[j].target
}

func findItinerary(tickets [][]string) []string {
	result := []string{}
	// map[出发机场] pair{目的地,是否被访问过}
	targets := make(map[string]pairs)
	// 记录映射关系
	for _, ticket := range tickets {
		if targets[ticket[0]] == nil { // 即map如果没有初始化，是不能赋值的。
			targets[ticket[0]] = make(pairs, 0) // 内存分配
		}
		targets[ticket[0]] = append(targets[ticket[0]], &pair{target: ticket[1], visited: false})
	}
	for k, _ := range targets { // 对目的地进行排序
		sort.Sort(targets[k])
	}
	result = append(result, "JFK") // 必须从“JFK”出发
	// 回溯函数定义
	var backTrack func() bool
	backTrack = func() bool {
		// base case, 结束条件
		if len(tickets)+1 == len(result) {
			return true
		}
		// 取出起飞航班对应的目的地
		for _, pair := range targets[result[len(result)-1]] {
			if pair.visited == false {
				result = append(result, pair.target) // 处理
				pair.visited = true                  // 处理
				if backTrack() {                     // 递归
					return true
				}
				result = result[:len(result)-1] // 回溯
				pair.visited = false            // 回溯
			}
		}
		return false
	}

	backTrack()

	return result
}

/*
以下参考官方和
https://leetcode.cn/problems/reconstruct-itinerary/solution/332-zhong-xin-an-pai-xing-cheng-shen-sou-hui-su-by/，
*/
func findItinerary(tickets [][]string) []string {
	var res []string
	m := map[string][]string{} // 每个起点对应的终点数组
	// 初始化，将图的邻接表存进字典
	for _, ticket := range tickets {
		m[ticket[0]] = append(m[ticket[0]], ticket[1])
	}
	// 按字典序排序
	for key := range m {
		sort.Strings(m[key]) // 把每个起点对应的数组按照字母排序
	}
	backTrack("JFK", m, &res)
	// 翻转数组
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
		// 去掉这条路
		m[cur] = m[cur][1:] // 注意,要在递归之前！
		backTrack(temp, m, res)

	}
	// 往下递归到终点之后开始入栈
	// 往上递归返回的入栈就是来时的路反着走回去，这样得到的数组是 路径的倒序
	// 越在链路后面的先入栈
	*res = append(*res, cur)
}