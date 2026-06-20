package lcr184

import "math"

/*
题目：设计自助结算系统
- https://leetcode.cn/problems/dui-lie-de-zui-da-zhi-lcof/description/?envType=study-plan-v2&envId=coding-interviews
- 请设计一个自助结账系统，该系统需要通过一个队列来模拟顾客通过购物车的结算过程，需要实现的功能有：
get_max()：获取结算商品中的最高价格，如果队列为空，则返回 -1
add(value)：将价格为 value 的商品加入待结算商品队列的尾部
remove()：移除第一个待结算的商品价格，如果队列为空，则返回 -1
注意，为保证该系统运转高效性，以上函数的均摊时间复杂度均为 O(1)

- 标签：队列，滑动窗口
- 解题思路：队列先进先出，滑动窗口
*/

type Checkout struct {
	queue []int
}

func Constructor() Checkout {
	return Checkout{}
}

func (this *Checkout) Get_max() int {
	if len(this.queue) == 0 {
		return -1
	}
	return max(this.queue)
}

func (this *Checkout) Add(value int) {
	this.queue = append(this.queue, value)
}

func (this *Checkout) Remove() int {
	if len(this.queue) == 0 {
		return -1
	}
	var res int
	res = this.queue[0]
	this.queue = this.queue[1:]
	return res
}

func max(temp []int) int {
	var res int = math.MinInt32
	for _, item := range temp {
		if res < item {
			res = item
		}
	}
	return res
}

/**
 * Your Checkout object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get_max();
 * obj.Add(value);
 * param_3 := obj.Remove();
 */
