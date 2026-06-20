package hot100

/*
35-146.LRU缓存
https://leetcode.cn/problems/lru-cache/?envType=study-plan-v2&envId=top-100-liked
*/

/*
关键点：
1. 以什么结构存储数据: 哈希表+双向循环链表
2. 以什么方式逐出最久未使用的关键字：链表的头节点就是最久未使用的
*/
type LRUCache struct {
	cap   int
	cache map[int]*DNode
	dummy *DNode // 虚拟头节点，前一个指针指向最后一个节点，后一个指针指向第一个节点
}

type DNode struct {
	key, val  int
	pre, next *DNode
}

func Constructor(capacity int) LRUCache {
	dummy := &DNode{}
	dummy.pre = dummy
	dummy.next = dummy
	return LRUCache{
		cap:   capacity,
		cache: map[int]*DNode{},
		dummy: dummy,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		// 将此节点移动最后
		// 先抽离此节点，删除此节点
		node.pre.next = node.next
		node.next.pre = node.pre
		// 移动最后，最后添加元素节点
		node.pre = this.dummy.pre
		node.next = this.dummy
		this.dummy.pre = node
		node.pre.next = node
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		// 将此节点移动最后
		// 先抽离此节点。删除此节点
		node.pre.next = node.next
		node.next.pre = node.pre
		// 最后添加元素节点
		node.pre = this.dummy.pre
		node.next = this.dummy
		this.dummy.pre = node
		node.pre.next = node
		// 更新值
		node.val = value
	} else {
		// 不存在，末尾添加新节点
		node := &DNode{
			key: key,
			val: value,
		}
		this.cache[key] = node
		// 最后添加元素
		node.pre = this.dummy.pre
		node.next = this.dummy
		this.dummy.pre = node
		node.pre.next = node
	}
	if len(this.cache) > this.cap {
		first := this.dummy.next
		delete(this.cache, first.key) // 这里用到key，所以node中要有key存储
		// 去掉第一个节点
		first.pre.next = first.next
		first.next.pre = first.pre
	}
}
