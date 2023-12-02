// 使用链表 + 哈希表(map) 实现 LRU 缓存淘汰算法
// 实现 3 个基本功能
// 1.缓存中查找一个数据
// 2.缓存中添加一个数据
// 3.缓存中删除一个数据
package combinedapplication

import "fmt"

type Node struct {
	key int
	value string
	pro *Node
	next *Node
}

func NewNode (key int, value string) *Node {
	return &Node {
		key : key, 
		value : value,
	}
}

type lruCache struct {
	nodeMap map[int]*Node
	size int
	capacity int
	// head 和 tail 作为哨兵，虚拟节点
	head *Node
	tail *Node
}

func initLruCache(capacity int) *lruCache {
	now := &lruCache {
		nodeMap : make(map[int]*Node),
		size : 0,
		capacity: capacity,
		head: NewNode(-1,""),
		tail: NewNode(-1,""),
	}
	now.head.pro = nil
	now.head.next = now.tail
	now.tail.pro = now.head
	now.tail.next = nil 
	return now
}

// 查找：缓存存在该值，返回该值，否则返回字符串 “not found!”
func (cache *lruCache) get(key int) string {
	if cache.size == 0 {
		return "cache is empty!"
	}
	nowNode, ok := cache.nodeMap[key]
	if !ok {
		return "not found!"
	}
	removeNode(nowNode)
	cache.addNodeAtHead(nowNode)
	return nowNode.value
}

//添加：
func (cache *lruCache) put (key int, value string) {
	if nowNode, ok :=  cache.nodeMap[key]; ok {
		nowNode.value = value
		removeNode(nowNode)
		cache.addNodeAtHead(nowNode)
		return
	}
	newnode := NewNode(key, value)
	if cache.size == cache.capacity {
		// 注意处理的先后顺序，如果先删除节点，再去删除 map 中数据的话，指针引用到错误节点上了
		//removeNode(cache.tail.pro)
		delete(cache.nodeMap, cache.tail.pro.key)
		removeNode(cache.tail.pro)
		cache.size --
		
		 for key, data := range cache.nodeMap {
			fmt.Println(key, data)
		 }
	}
	cache.addNodeAtHead(newnode)
	cache.nodeMap[key] = newnode
	cache.size ++
}

func (cache *lruCache) remove(key int) string {
	if nownode, ok := cache.nodeMap[key]; ok {
		delete(cache.nodeMap, key)
		removeNode(nownode)
		return nownode.value
	}
	return "缓存中无法找到该值"
}

func removeNode(node *Node) {
	node.next.pro = node.pro
	node.pro.next = node.next
}

func (cache *lruCache) addNodeAtHead(node *Node) {
	node.next = cache.head.next
	cache.head.next.pro = node
	cache.head.next = node
	node.pro = cache.head
}



