package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	head := &ListNode{Val:1, Next:&ListNode{Val:2, Next:&ListNode{Val:3, Next:&ListNode{Val:4, Next:&ListNode{Val:5}}}}}
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

}

func print(head *ListNode) {
	for head != nil {
		fmt.Printf("%v->", head.Val)
		head = head.Next
	}
	fmt.Println()
}
