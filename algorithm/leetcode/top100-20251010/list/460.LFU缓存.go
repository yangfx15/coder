package main

import "fmt"

//请你为 最不经常使用（LFU）缓存算法设计并实现数据结构。
//
//实现 LFUCache 类：
//
//LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
//int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
//void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。当缓存达到其容量 capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最久未使用 的键。
//为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。
//
//当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。
//
//函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
//示例：
//
//输入：
//["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]
//输出：
//[null, null, null, 1, null, -1, 3, null, -1, 3, 4]
//
//解释：
//// cnt(x) = 键 x 的使用计数
//// cache=[] 将显示最后一次使用的顺序（最左边的元素是最近的）
//LFUCache lfu = new LFUCache(2);
//lfu.put(1, 1);   // cache=[1,_], cnt(1)=1
//lfu.put(2, 2);   // cache=[2,1], cnt(2)=1, cnt(1)=1
//lfu.get(1);      // 返回 1
//                 // cache=[1,2], cnt(2)=1, cnt(1)=2
//lfu.put(3, 3);   // 去除键 2 ，因为 cnt(2)=1 ，使用计数最小
//                 // cache=[3,1], cnt(3)=1, cnt(1)=2
//lfu.get(2);      // 返回 -1（未找到）
//lfu.get(3);      // 返回 3
//                 // cache=[3,1], cnt(3)=2, cnt(1)=2
//lfu.put(4, 4);   // 去除键 1 ，1 和 3 的 cnt 相同，但 1 最久未使用
//                 // cache=[4,3], cnt(4)=1, cnt(3)=2
//lfu.get(1);      // 返回 -1（未找到）
//lfu.get(3);      // 返回 3
//                 // cache=[3,4], cnt(4)=1, cnt(3)=3
//lfu.get(4);      // 返回 4
//                 // cache=[3,4], cnt(4)=2, cnt(3)=3

type LFUNode struct {
	Key, Value, Freq int
	Prev, Next, Tail *LFUNode
}

type LFUCache struct {
	Record             map[int]*LFUNode // 数据记录
	FreqQue            map[int]*LFUNode // 频率记录
	FreqCount          map[int]int      // 频率数量
	Cap, Size, minFreq int
}

func LFUConstructor(capacity int) LFUCache {
	lru := LFUCache{
		Record:    make(map[int]*LFUNode, 0),
		FreqQue:   make(map[int]*LFUNode, 0),
		FreqCount: make(map[int]int, 0),
		Cap:       capacity,
	}

	return lru
}

func (this *LFUCache) UpdateFreq(node *LFUNode) {
	if this.FreqCount[node.Freq-1] == 1 {
		this.FreqQue[node.Freq-1] = nil // 删除

		this.FreqCount[node.Freq-1]--
		fmt.Printf("remove freq_queue %v, minFreq: %v, queue_num: %v\n", node.Freq-1, this.minFreq, this.FreqCount[node.Freq-1])
		this.minFreq = node.Freq
	} else if node.Freq-1 > 0 {
		// 删除该频率队列里的元素
		freqNode := this.FreqQue[node.Freq-1].Next
		for freqNode != nil && freqNode.Key != node.Key {
			freqNode = freqNode.Next
		}
		freqNode.Prev.Next = freqNode.Next
		freqNode.Next.Prev = freqNode.Prev

		this.FreqCount[node.Freq-1]--
		fmt.Printf("remove %v from freq_queue %v, queue_num: %v\n", node.Key, node.Freq-1, this.FreqCount[node.Freq-1])
	}

	// 初始化
	if this.FreqCount[node.Freq] == 0 {
		head, tail := &LFUNode{}, &LFUNode{}
		this.FreqQue[node.Freq] = head
		this.FreqQue[node.Freq].Tail = tail
		this.FreqQue[node.Freq].Next = this.FreqQue[node.Freq].Tail
		this.FreqQue[node.Freq].Tail.Prev = this.FreqQue[node.Freq]

		if this.minFreq == 0 || node.Freq < this.minFreq {
			this.minFreq = node.Freq
		}

		fmt.Printf("init %v freq to queue, minFreq: %v\n", node.Freq, this.minFreq)
	}

	// 加到第一个节点位置
	this.FreqCount[node.Freq]++
	this.AddBeforeNode(node, this.FreqQue[node.Freq], this.FreqQue[node.Freq].Next)

	fmt.Printf("move %v to freq %v first\n", node.Key, node.Freq)
}

func (this *LFUCache) Get(key int) int {
	fmt.Printf("----------------- Get %v ----------------\n", key)
	node, ok := this.Record[key]
	if !ok {
		return -1
	}

	node.Freq++
	this.UpdateFreq(node)

	return node.Value
}

func (this *LFUCache) AddBeforeNode(node, before, after *LFUNode) {
	node.Prev, node.Next = before, after
	before.Next, after.Prev = node, node
}

func (this *LFUCache) Put(key int, value int) {
	fmt.Printf("----------------- Put %v ----------------\n", key)
	node, ok := this.Record[key]
	if ok {
		// 存在即更新
		node.Value = value
		node.Freq++
		this.UpdateFreq(node)

		return
	}
	// 新增
	if this.Cap == this.Size {
		delQueue := this.FreqQue[this.minFreq]
		tail := delQueue.Tail.Prev

		if this.FreqCount[this.minFreq] == 1 {
			delQueue = nil
		} else {
			tail.Prev.Next = tail.Next
			tail.Next = tail.Prev
		}

		delete(this.Record, tail.Key)
		fmt.Printf("remove record %v\n", tail.Key)
		this.Size--
	}

	this.Size++
	newNode := &LFUNode{Key: key, Value: value, Freq: 1}
	fmt.Printf("key %v add to %v\n", key, value)
	this.Record[key] = newNode
	this.UpdateFreq(newNode)
}

func main() {
	//["LFUCache","put","put","get","get","get","put","put","get","get","get","get"]
	//[[3],[2,2],[1,1],[2],[1],[2],[3,3],[4,4],[3],[2],[1],[4]]
	//预期：[null,null,null,2,1,2,null,null,-1,2,1,4]
	obj := LFUConstructor(3)
	obj.Put(2, 2)
	obj.Put(1, 1)
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(2))
	obj.Put(3, 3)
	obj.Put(4, 4)
	fmt.Println(obj.Get(3))
	fmt.Println(obj.Get(2))
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(4))
}
