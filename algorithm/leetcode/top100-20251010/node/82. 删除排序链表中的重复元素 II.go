package main

//给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
//
//示例 1：
//
//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
//示例 2：
//
//输入：head = [1,1,1,2,3]
//输出：[2,3]
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := &ListNode{Next: head}
	hair := prev
	for head != nil {
		if head.Next == nil {
			return hair.Next
		}
		next := head.Next
		for next != nil && next.Val == head.Val {
			next = next.Next
		}
		if next != head.Next {
			prev.Next = next
			head = next
			continue
		}
		prev = prev.Next
		head = head.Next
	}
	return hair.Next
}

func main()  {
	// [1,2,3,3,4,4,5]
	// [1,1,1,2,3]
	printList(deleteDuplicates(&ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 3,
					Next: &ListNode{Val: 4,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 5}}}}}}}))
	printList(deleteDuplicates(&ListNode{Val: 1,
		Next: &ListNode{Val: 1,
			Next: &ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3}}}}}))
}