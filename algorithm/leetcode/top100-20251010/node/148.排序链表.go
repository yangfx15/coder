package main

// 给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
//
//示例 1：
//
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
//示例 2：
//
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]
//示例 3：
//
//输入：head = []
//输出：[]

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	next := slow.Next
	slow.Next = nil
	return sortTwoList(sortList(head), sortList(next))
}

func sortTwoList(a, b *ListNode) *ListNode {
	pre := new(ListNode)
	hair := pre

	for a != nil && b != nil {
		if a.Val < b.Val {
			pre.Next = a
			a = a.Next
		} else {
			pre.Next = b
			b = b.Next
		}
		pre = pre.Next
	}

	if a != nil {
		pre.Next = a
	}

	if b != nil {
		pre.Next = b
	}

	return hair.Next
}
