package lru2

/*
LRUCache
使用双向循环链表实现
leetcode 版
*/
type LRUCache struct {
	len  int
	cap  int
	hmap map[int]*Node // 哈希表
	list *Node         // 双向循环链表
}

// Node 双向链表的节点
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

func (cache *LRUCache) Get(key int) int {
	n, ok := cache.hmap[key]
	if !ok {
		return -1 // 没有找到
	}
	// 找到，就要移动到头部
	cache.moveToHead(n)
	return n.value
}

func (cache *LRUCache) Put(key int, value int) {
	n, ok := cache.hmap[key]
	if ok {
		n.value = value
		cache.moveToHead(n)
		return
	}

	if cache.len == cache.cap {
		n = cache.list.prev
		delete(cache.hmap, n.key)
	} else {
		n = &Node{}
		cache.len++
	}

	n.key = key
	n.value = value
	cache.hmap[key] = n
	cache.moveToHead(n)

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
