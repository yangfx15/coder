package _43

// 重排链表
// 1->2->3->4->5
// 1->5->2->4->3

type Node struct {
	Value int
	Next  *Node
}

func revNode(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	// 到中间断开，重连
	mid := head
	for head.Next != nil && head.Next.Next != nil {
		head = head.Next.Next
		mid = mid.Next
	}

	mid = revCur(mid)

}

func mergeTwoNode(n1, n2 *Node) *Node {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}

	re := &Node{}
	for n1 != nil && n2 != nil {
		if n1.Next != nil {
			re.Next = n1
			n1 = n1.Next
		}
	}
}

func revCur(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	var prev *Node
	cur := head
	for cur != nil {
		next := cur.Next
		prev = cur
		cur = next
	}

	return prev
}
