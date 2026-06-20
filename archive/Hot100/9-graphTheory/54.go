package hot100

/*
54-208.实现Trie（前缀树）
https://leetcode.cn/problems/implement-trie-prefix-tree/?envType=study-plan-v2&envId=top-100-liked
*/
/*
关键点：
1. 前缀树节点如何存储：因为都是小写字母，所以使用26个字母存储。
2. 如何从前缀树根节点遍历：循环深度
*/
// type Trie struct {
// 	child [26]*Trie
// 	isEnd bool // 是否是某个词的结尾节点
// }

// func Constructor() Trie {
// 	return Trie{}
// }

// func (this *Trie) Insert(word string) {
// 	node := this // 从根节点开始
// 	for _, item := range word {
// 		num := item - 'a'
// 		if node.child[num] == nil {
// 			node.child[num] = &Trie{}
// 		}
// 		node = node.child[num]
// 	}
// 	node.isEnd = true // 到词结尾表示
// }

// func (this *Trie) searchPrefix(word string) *Trie {
// 	node := this
// 	for _, item := range word {
// 		num := item - 'a'
// 		if node.child[num] == nil {
// 			return nil
// 		}
// 		node = node.child[num]
// 	}
// 	return node
// }

// func (this *Trie) Search(word string) bool {
// 	node := this.searchPrefix(word)
// 	if node == nil {
// 		return false
// 	}
// 	return node.isEnd
// }

// func (this *Trie) StartsWith(prefix string) bool {
// 	node := this.searchPrefix(prefix)
// 	if node == nil {
// 		return false
// 	}
// 	return true
// }

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

/*
思路：使用map存储子节点
*/
type Trie struct {
	children map[rune]*Trie
	isEnd    bool // 是一个词结尾
}

func Constructor() Trie {
	return Trie{children: make(map[rune]*Trie)} // 注意这里初始化
}

func (this *Trie) Insert(word string) {
	node := this
	for _, item := range word {
		if node.children[item] == nil {
			node.children[item] = &Trie{children: make(map[rune]*Trie)}
		}
		node = node.children[item]
	}
	node.isEnd = true
}

func (this *Trie) searchPrefix(word string) *Trie {
	node := this
	for _, item := range word {
		if node.children[item] == nil {
			return nil
		}
		node = node.children[item]
	}
	return node
}

func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	if node == nil {
		return false
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.searchPrefix(prefix)
	if node == nil {
		return false
	}
	return true
}
