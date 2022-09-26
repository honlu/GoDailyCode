/* 使用双向循环链表实现
leetcode版
*/
type LRUCache struct {
	len  int
	cap  int
	hmap map[int]*Node // 哈希表
	list *Node         // 双向循环链表
}

// 双向链表的节点
type Node struct {
	key, value int
	prev, next *Node
}

func Constructor(capacity int) LRUCache {
	list := &Node{}
	list.next = list // 循环链表
	list.prev = list

	l := LRUCache{
		len:  0,
		cap:  capacity,
		hmap: make(map[int]*Node, capacity),
		list: list,
	}
	return l
}

// 节点移动到头部
func (cache *LRUCache) moveToHead(n *Node) {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}

	n.next = cache.list.next
	n.prev = cache.list
	n.next.prev = n
	n.prev.next = n
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.hmap[key]
	if !ok {
		return -1 // 没有找到
	}
	// 找到，就要移动到头部
	this.moveToHead(n)
	return n.value
}

func (this *LRUCache) Put(key int, value int) {
	n, ok := this.hmap[key]
	if ok {
		n.value = value
		this.moveToHead(n)
		return
	}

	if this.len == this.cap {
		n = this.list.prev
		delete(this.hmap, n.key)
	} else {
		n = &Node{}
		this.len++
	}

	n.key = key
	n.value = value
	this.hmap[key] = n
	this.moveToHead(n)

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */