package main

import "fmt"

type Node struct {
	index, value int
}

type LRUCache struct {
	lru   map[int]*Node // 存储key ,value
	size  int           // 约束容器数
	queue []int         // 移动成本会比较高
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		size:  capacity,
		lru:   make(map[int]*Node),
		queue: []int{},
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.lru[key]
	if !ok {
		return -1
	}
	// 存在更新位置
	index := node.index
	this.queue = append(this.queue[:index], this.queue[index+1:]...)
	this.queue = append(this.queue, key)
	// 刷新index
	this.refreshIndex()
	return node.value
}

func (this *LRUCache) refreshIndex() {
	for i, key := range this.queue {
		this.lru[key].index = i
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		node, _ := this.lru[key]
		node.value = value // 为什么这里没有更新呢？
		return
	}
	if len(this.queue) == this.size {
		// 删除最久的然后添加到末尾
		delete(this.lru, this.queue[0])
		this.queue = this.queue[1:]
		this.refreshIndex() // 注意这里要刷新

	}
	this.queue = append(this.queue, key)
	this.lru[key] = &Node{
		index: len(this.queue) - 1,
		value: value,
	}
}

/*
*
  - Your LRUCache object will be instantiated and called as such:
  - obj := Constructor(capacity);
  - param_1 := obj.Get(key);
  - obj.Put(key,value);

最近最少使用：1. 最近使用元素要插入到最后，2. 队列满了之后要删除历史最久的.
map+slice思路：
Get: O(n)
Put: O(n)

因为 slice 删除和刷新 index 都需要移动/遍历。「题目超时啦」
后面优化
*/
func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	obj.Put(3, 3)
	fmt.Println(obj.Get(2))
	obj.Put(4, 4)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(4))
	obj.Put(4, 5)
	fmt.Println(obj.Get(4))
}
