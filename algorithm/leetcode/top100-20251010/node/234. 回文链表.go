package main

import "fmt"

//给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//
//示例 1：
//
//输入：head = [1,2,2,1]
//输出：true
//示例 2：
//
//输入：head = [1,2]
//输出：false

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	next := slow.Next
	slow.Next = nil
	// 反转前半个链表

	//printList(head)
	prev := reverseSlowNode(head)
	//printList(prev)
	//printList(next)

	return checkPalindrome(prev, next) || checkPalindrome(prev.Next, next)
}

func checkPalindrome(prev, next *ListNode) bool {
	for next != nil && prev != nil {
		if next.Val != prev.Val {
			return false
		}
		next, prev = next.Next, prev.Next
	}

	return next == nil && prev == nil
}

func reverseSlowNode(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func main()  {
	fmt.Println(isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}))
	fmt.Println(isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}))
	fmt.Println(isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 0}}))
}