/*
2022-9-26

LRU: Least Recently Used 最近最少使用算法。
是一种内存数据淘汰策略，常见使用场景是当内存不足时，需要淘汰最近最少使用的数据。
LRU常用语缓存系统的淘汰策略。

使用双向循环链表实现LRU
，这里可能并不是并发安全的哈，因为没有使用锁
*/
package lru

import "fmt"

// Node 链表的节点
type Node struct {
	prev, next *Node

	list *LRU

	key   string
	value interface{}
}

// LRU 缓存
type LRU struct {
	root *Node // 根节点
	cap  int   // 当前缓存容量
	len  int   // 缓存的长度
}

// NewLRU 新建一个LRU缓存
func NewLRU(cap int) *LRU {
	l := &LRU{
		root: &Node{},
		cap:  cap,
	}
	l.root.prev = l.root
	l.root.next = l.root
	l.root.list = l
	return l
}

// LRU的方法：Get 获取缓存数据
// 如果获取到数据，就把这个节点移动到链表的头部。即更新节点
// 如果没有获取到，就返回nil
func (l *LRU) Get(key string) interface{} {
	defer l.debug() // 打印LRU
	n := l.get(key)
	if n == nil {
		return nil
	}
	return n.value
}
func (l *LRU) get(key string) *Node {
	for n := l.root.next; n != l.root; n = n.next { // 循环遍历双链表
		if n.key == key { // 找到这个元素节点，要进行更新，移动到链表的头部
			n.prev.next = n.next // 首先在链表中移除这个节点
			n.next.prev = n.next

			// 然后更新到链表的头部
			n.next = l.root.next
			l.root.next.prev = n
			l.root.next = n
			n.prev = l.root
			return n
		}
	}
	return nil // 没有找到这个节点，返回nil
}

// Put 写入缓存数据
// 如果key已经存在，那么更新值
// 如果key 不存在，那么插入到第一个节点
// 当缓存容量满了的时候，会自动删除最后的数据
func (l *LRU) Put(key string, value interface{}) {
	defer l.debug() // 打印链表
	n := l.get(key)
	if n != nil {
		n.value = value // 更新节点的值
		return
	}

	// 缓存满了,首先删除最后的数据
	if l.len == l.cap {
		// 首先删除最后的节点数据
		last := l.root.prev     // 获取最后一个节点
		last.prev.next = l.root // 然后删除
		l.root.prev = last.prev

		last.list = nil
		last.prev = nil
		last.next = nil
		l.len--
	}
	// 新建节点插入到最后
	node := &Node{
		key:   key,
		value: value,
	}
	head := l.root.next
	head.prev = node
	node.next = head
	node.prev = l.root
	l.root.next = node
	l.len++
	node.list = l // ?
}

// debug for LRU
func (l *LRU) debug() {
	fmt.Println("lru len: ", l.len)
	fmt.Println("lru cap: ", l.cap)
	for n := l.root.next; n != l.root; n = n.next {
		fmt.Printf("%s:%v->", n.key, n.value)
	}
	fmt.Println()
}
