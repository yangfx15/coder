package main

//设计你的循环队列实现。 循环队列是一种线性数据结构，其操作表现基于 FIFO（先进先出）原则并且队尾被连接在队首之后以
//形成一个循环。它也被称为“环形缓冲器”。
//
//循环队列的一个好处是我们可以利用这个队列之前用过的空间。在一个普通队列里，一旦一个队列满了，我们就不能插入
//下一个元素，即使在队列前面仍有空间。但是使用循环队列，我们能使用这些空间去存储新的值。
//
//你的实现应该支持如下操作：
//
//MyCircularQueue(k): 构造器，设置队列长度为 k 。
//Front: 从队首获取元素。如果队列为空，返回 -1 。
//Rear: 获取队尾元素。如果队列为空，返回 -1 。
//enQueue(value): 向循环队列插入一个元素。如果成功插入则返回真。
//deQueue(): 从循环队列中删除一个元素。如果成功删除则返回真。
//isEmpty(): 检查循环队列是否为空。
//isFull(): 检查循环队列是否已满。

type MyCircularQueue struct {
	Queue      []int
	Start, End int
	Cap, Size  int
}

func QueueConstructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Cap:   k,
		Queue: make([]int, k),
	}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.Cap == q.Size {
		return false
	}
	q.Size++
	q.Queue[q.End] = value

	// 循环指定
	if q.End == q.Cap-1 {
		q.End = 0
	} else {
		q.End++
	}

	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	// 删除队头元素
	if q.Size == 0 {
		return false
	}

	q.Size--
	if q.Start == q.Cap-1 {
		q.Start = 0
	} else {
		q.Start++
	}

	return true
}

func (q *MyCircularQueue) Front() int {
	if q.Size == 0 {
		return -1
	}

	return q.Queue[q.Start]
}

func (q *MyCircularQueue) Rear() int {
	if q.Size == 0 {
		return -1
	}

	if q.End == 0 {
		return q.Queue[q.Cap-1]
	}
	return q.Queue[q.End-1]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.Size == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.Size == q.Cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
