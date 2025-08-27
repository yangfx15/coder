package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode{Val:1, Next:&ListNode{Val:2, Next:&ListNode{Val:3, 
		Next:&ListNode{Val:4, Next:&ListNode{Val:5, Next:&ListNode{Val:6}}}}}}
	res := reverseList(head)
	for res != nil {
		fmt.Printf("%v->", res.Val)
		res = res.Next
	}
}

// input: 1->2->3->4->5
// output: 1->5->2->4->3
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow, first := head, head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	next := slow.Next
	slow.Next = nil

	// 反转链表
	next = reverseNode(next)

	// 1->2->3
	// 5->4
	// res := first
	// for first != nil && next != nil {
	// 	h1Temp, h2Temp := first.Next, next.Next  // 2,4
	// 	first.Next = next
	// 	next.Next = h1Temp

	// 	first = h1Temp
	// 	next = h2Temp
	// }

	return mergeTwoList(first, next)
}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%v->", head.Val)
		head = head.Next
	}
	fmt.Println()
}

// 1->2->3
func reverseNode(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func mergeTwoList(h1, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}

	res := &ListNode{}
	pre := res
	for h1 != nil && h2 != nil {
		if res.Next == nil {
			res.Next = &ListNode
		}
		res = res.Next
		res.val = h1.Val
		res = res.Next
		h1 = h1.Next
		res = &ListNode{Val: h2.Val, Next:&ListNode{}}
		h2 = h2.Next
	}

	if h1 != nil {
		res.Next = h1
	}
	if h2 != nil {
		res.Next = h2
	}
	return pre.Next
}