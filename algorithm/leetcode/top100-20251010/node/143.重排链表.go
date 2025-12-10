package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// input: 1->2->3->4->5
// output: 1->5->2->4->3
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	next := slow.Next
	slow.Next = nil
	next = reverse(next)
	//print(next)
	//print(head)

	mergeTwoList(head, next)
}

func mergeTwoList(a, b *ListNode) {
	for a != nil && b != nil {
		aTemp, bTemp := a.Next, b.Next
		a.Next, a = b, aTemp
		b.Next, b = a, bTemp
	}
}

func reverse(header *ListNode) *ListNode {
	var pre *ListNode
	root := header
	for root != nil {
		next := root.Next
		root.Next = pre
		pre = root
		root = next
	}
	return pre
}

//func print(head *ListNode) {
//	cur := head
//	for cur != nil {
//		fmt.Printf("%v->", cur.Val)
//		cur = cur.Next
//	}
//	fmt.Println()
//}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3,
		Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}
	reorderList(head)
	for head != nil {
		fmt.Printf("%v->", head.Val)
		head = head.Next
	}
}
