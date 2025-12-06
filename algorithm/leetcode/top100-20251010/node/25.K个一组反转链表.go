package main

// 给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。
//k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//示例 1：
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]

//示例 2：
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]

//提示：
//链表中的节点数目为 n
//1 <= k <= n <= 5000
//0 <= Node.val <= 1000

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 1.节点交换
	// 2.剩余节点保持原顺序 -> 循环遍历交换节点后，连接后续节点

	// 设置头结点返回
	pre := new(ListNode)
	pre.Next = head
	hair := pre
	for head != nil {
		end := head
		// 循环节点头结点
		for i := 1; i < k; i++ {
			if end.Next == nil {
				return hair.Next
			}
			end = end.Next
		}
		next := end.Next
		end.Next = nil

		// 反转k个节点
		//printList(head)
		head, end = reverseKList(head, end)
		//printList(head)

		end.Next = next
		pre.Next = head

		head = next
		pre = end
	}

	return hair.Next
}

func reverseKList(a, b *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	root := a
	for pre != b {
		next := root.Next
		root.Next = pre
		pre = root
		root = next
	}
	return pre, a
}

func main() {
	printList(reverseKGroup(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 5, Next: nil}},
			}},
	}, 2))

	printList(reverseKGroup(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 5, Next: nil}},
			}},
	}, 1))

	printList(reverseKGroup(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 5, Next: &ListNode{
							Val:  6,
							Next: nil,
						}}},
			}},
	}, 2))

	printList(reverseKGroup(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2, Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4, Next: &ListNode{
						Val: 5, Next: &ListNode{
							Val:  6,
							Next: nil,
						}}},
			}},
	}, 3))
}
