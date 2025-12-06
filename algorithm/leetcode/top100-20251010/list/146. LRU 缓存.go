package main

import "fmt"

//请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
//实现 LRUCache 类：
//LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。
//如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

//示例：
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4

type Node struct {
	Key, Value int
	Next       *Node
	Prev       *Node
}

type LRUCache struct {
	FirstNode *Node
	TailNode  *Node
	Record    map[int]*Node // 缓存记录元素
	Cap       int           // 总容量
	Size      int           // 当前长度
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Record:    make(map[int]*Node, 0),
		FirstNode: &Node{}, // 虚拟节点，防止空指针异常
		TailNode:  &Node{}, // 虚拟节点，防止空指针异常
		Cap:       capacity,
	}
	lru.FirstNode.Prev, lru.FirstNode.Next = lru.TailNode, lru.TailNode
	lru.TailNode.Next, lru.TailNode.Prev = lru.FirstNode, lru.FirstNode

	return lru
}

func (this *LRUCache) MoveToHead(node *Node) {
	node.Next, node.Prev = this.FirstNode.Next, this.FirstNode
	this.FirstNode.Next.Prev, this.FirstNode.Next = node, node
}

func (this *LRUCache) Remove(node *Node) {
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Record[key]
	if !ok {
		return -1
	}
	this.Remove(node)
	this.MoveToHead(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Record[key]
	if ok {
		// 存在则更新
		node.Value = value
		this.Remove(node)
		this.MoveToHead(node)
		//fmt.Printf("key %v update to %v\n", key, value)

		return
	}
	// 若队列已满，则清除尾结点
	if this.Cap == this.Size {
		tail := this.TailNode.Prev
		this.Remove(tail)
		delete(this.Record, tail.Key)
		//fmt.Printf("key %v remove %v\n", tail.Key, tail.Value)
		this.Size--
	}

	this.Size++
	newNode := &Node{Key: key, Value: value}
	this.Record[key] = newNode

	this.MoveToHead(newNode)

	//fmt.Printf("key %v add %v\n", key, value)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
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
}
