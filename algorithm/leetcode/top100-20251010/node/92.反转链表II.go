package main

// 给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
//
//示例 1：
//
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//示例 2：
//
//输入：head = [5], left = 1, right = 1
//输出：[5]

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if left > right || left == right {
		return head
	}

	// 先找 left->right 的节点做反转
	// 然后拼接所有节点

	pre := new(ListNode)
	pre.Next = head
	hair := pre

	for i := 1; i < left; i++ {
		if pre.Next == nil {
			return hair.Next
		}
		pre = pre.Next
	}

	end := pre.Next
	start := end
	pre.Next = nil

	for i := 0; i < right-left; i++ {
		if end.Next != nil {
			end = end.Next
		}
	}

	next := end.Next
	end.Next = nil

	start, end = reverseKNode(start, end)

	pre.Next = start
	end.Next = next

	return hair.Next
}

func reverseKNode(start, end *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	head := start
	for pre != end {
		next := head.Next
		head.Next = pre

		pre = head
		head = next
	}
	return pre, start
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 5, Next: nil}},
			}},
	}

	PrintNode(reverseBetween(head, 2, 6))
}
