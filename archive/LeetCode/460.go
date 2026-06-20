/* 
480. LFU缓存
2022-12-27
link: https://leetcode.cn/problems/lfu-cache/
question:
	为最不经常使用（LFU）缓存算法设计并实现的数据结构。
	实现 LFUCache 类：
		LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象。
		int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
		void put(int key, int value) - 如果键 key 已存在，则变更其值；
		如果键不存在，请插入键值对。当缓存达到其容量 capacity 时，则应该在插入新项之前，移除最不经常使用的项。
		在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
	为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
answer:
	难点get和put平均时间复杂度都是O(1)
	参考：https://leetcode.cn/problems/lfu-cache/solution/by-wonbin-d06u/
	使用container包中的双向链表和map实现：
	两个哈希表：
		1.map:key为添加进去的key,value为双向链表中的元素即Node类型
		2.map：key为频率值，value为计数值对应搜有Node类型的双向链表
*/
package main

import "container/list" // 列表，内部实现原理是双链表，可以高效进行任意位置的元素插入和删除操作。

type LFUCache struct {
	hash map[int]*list.Element // key，value是链表的元素类型，封装key、value、freq(计数器)
	cap int
}

type Node struct{
	key int
	value int
	freq int  // 计数器
}

var minFreq = 0

var freHash map[int]*list.List  // 定义一个哈希：key为使用计数器值，值为双向链表

func Constructor(capacity int) LFUCache {
	if capacity < 0 || capacity > 1e4{ // 不满足提示条件的
		return LFUCache{}
	}
	// 正式构建
	cache := LFUCache{
		hash: make(map[int]*list.Element),
		cap: capacity,
	}
	freHash = make(map[int]*list.List)

	return cache
}


func (this *LFUCache) Get(key int) int {
	// 1. 查询缓存，没有返回-1
	var exist bool
	var temp *list.Element

	if temp, exist = this.hash[key];!exist{
		return -1
	} 
	// 2. 如果存在，返回结果，freq+1, 更新minFreq，freqHash,freqList
	node := temp.Value.(Node)  // 这里什么意思？使用类型断言断言专程具体类型，即interface{}转成Nodel类型
	oldFreq := node.freq
	node.freq++

	var freqList *list.List   // 双向链表
	freqList = freHash[oldFreq]  // 此双向链表一定存在
	freqList.Remove(temp)  // 在旧计数值对应的链表中删除temp元素
	if freqList.Len() == 0{ // 当旧的计数器值对应的双向链表没有元素后
		if minFreq == oldFreq{ // 并且是最小计数的值时，要更新计数最小的值
			minFreq = node.freq
		}
	}
	// 再根据新计算的值（使用频率值）添加到相应的双链表中
	if freqList, exist = freHash[node.freq]; !exist{ // 如果新计数的值，不存在双向链表：要新建双向链表
		freqList = list.New()
	}
	this.hash[key] = freqList.PushFront(node)  // 双向链表：从队列头插入元素，PushFront会返回一个*list.Elemente结构
	// 如果以后需要删除插入的元素，则只通过 *list.Element 配合 Remove() 方法就可删除，可以让删除更加效率化，也是双链表特性之一。
	freHash[node.freq] = freqList

	return node.value
}


func (this *LFUCache) Put(key int, value int)  {
	// 不符合条件的添加
	if this.cap == 0{
		return
	}
	if value < 0 && value > 1e9{
		return
	}
	if key < 0 && key > 1e5{
		return 
	}
	// 符号条件的元素添加
	var temp *list.Element
	var node Node
	var exist bool
	var freqList *list.List

	if temp, exist := this.hash[key]; exist{ // 查看要添加的元素是已经存在
		// 1. 存在就更新计数的值
		node = temp.Value.(Node)  // 类型断言转换成具体Node类型
		node.value = value  // 注意key对应的value可能更新了
		oldFreq := node.freq
		node.freq++

		// 更新freHash,freList
		freqList = freHash[oldFreq]
		freqList.Remove(temp) // list删除Element
		if freqList.Len() == 0{ // 删除后list长度为0
			if minFreq == oldFreq{
				minFreq = node.freq
			}
		}
		// freq加1后对应的双链表
		if freqList, exist = freHash[node.freq]; !exist{
			freqList = list.New()
		}
		this.hash[key] = freqList.PushFront(node)  // PushFront从列表头插入一个元素
		freHash[node.freq] = freqList
		
		return
	}

	// 2. 不存在，如果缓存满了，查找最小使用的minFreq对应的key，然后删掉。
	// 如果存在多个最小使用的Key，删掉最旧的那个()
	var minList *list.List
	if this.cap == len(this.hash){
		minList = freHash[minFreq] 
		//注意前面元素插入列表都是从头部插入的，所以前进先出原则即前进的都是相对来说旧的，所以从尾部选择删除
		temp = minList.Back()  // 最久最小使用频次的.Back()返回列表的最后一个元素或空值nil
		minList.Remove(temp) // list中删除
		node = temp.Value.(Node)

		delete(this.hash, node.key) // map中删除
	}

	// 3. 缓存还没有满，直接放入缓存，更新freq
	node = Node{key, value , 1} // 这里终于揭秘了
	minFreq = node.freq
	
	if freqList, exist = freHash[minFreq]; !exist{
		freqList = list.New()
	}

	this.hash[key] = freqList.PushFront(node) // map中添加
	freHash[minFreq] = freqList
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */