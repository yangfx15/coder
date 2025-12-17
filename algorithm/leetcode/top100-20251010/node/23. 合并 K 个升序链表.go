package main

// 给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
// 输出：[1,1,2,3,4,4,5,6]
// 解释：链表数组如下：
// [
//
//	1->4->5,
//	1->3->4,
//	2->6
//
// ]
// 将它们合并到一个有序链表中得到。
// 1->1->2->3->4->4->5->6
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}
	mid := len(lists) / 2
	return mergeTwoLists(mergeKLists(lists[:mid]), mergeKLists(lists[mid:]))
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	prev := head
	for l1 != nil && l2 != nil {
		if head.Next == nil {
			head.Next = &ListNode{}
			head = head.Next
		}
		if l1.Val <= l2.Val {
			head.Val = l1.Val
			l1 = l1.Next
		} else {
			head.Val = l2.Val
			l2 = l2.Next
		}
	}

	if l1 != nil {
		head.Next = l1
	}
	if l2 != nil {
		head.Next = l2
	}

	return prev.Next
}
