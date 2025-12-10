package main

// 给定单链表的头节点 head ，将所有索引为奇数的节点和索引为偶数的节点分别分组，保持它们原有的相对顺序，然后把偶数索引节点分组连接到奇数索引节点分组之后，返回重新排序的链表。
//
//第一个节点的索引被认为是 奇数 ， 第二个节点的索引为 偶数 ，以此类推。
//
//请注意，偶数组和奇数组内部的相对顺序应该与输入时保持一致。
//
//你必须在 O(1) 的额外空间复杂度和 O(n) 的时间复杂度下解决这个问题。
//
//示例 1:
//
//输入: head = [1,2,3,4,5]
//输出: [1,3,5,2,4]
//示例 2:
//
//输入: head = [2,1,3,5,6,4,7]
//输出: [2,3,6,7,1,5,4]
//
//提示:
//
//n ==  链表中的节点数
//0 <= n <= 104
//-106 <= Node.val <= 106

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 一个节点连接奇数，一个节点连接偶数，再尾首相连
	odd, even := head, head.Next
	hair, next := odd, even
	for even != nil && even.Next != nil {
		odd2 := even.Next
		odd.Next = odd2

		even2 := odd2.Next
		even.Next = even2

		odd = odd.Next
		even = even.Next
	}
	odd.Next = next

	return hair
}

func main() {
	PrintNode(oddEvenList(Node1))
	PrintNode(oddEvenList(Node2))
}
