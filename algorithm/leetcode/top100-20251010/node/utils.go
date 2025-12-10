package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Node1 = [1,2,3,4,5]
var Node1 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: &ListNode{Val: 5},
			},
		},
	},
}

// Node2 = [2,1,3,5,6,4,7]
var Node2 = &ListNode{
	Val: 2,
	Next: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{Val: 7},
					},
				},
			},
		},
	},
}

func PrintNode(head *ListNode) {
	for head != nil {
		fmt.Printf("%v->", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func printList(h *ListNode) {
	for h != nil {
		fmt.Printf("%v->", h.Val)
		h = h.Next
	}
	fmt.Println()
}
