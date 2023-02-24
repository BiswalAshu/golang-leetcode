package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	res := head
	if head == nil {
		return head
	}
	if head.Next == nil && head.Val == val {
		return head.Next
	}
	// node := &ListNode{Val: val, Next: nil}
	for head != nil {
		if head.Next.Val == val {
			head.Next = head.Next.Next
		} else {
			head = head.Next
		}

	}
	return res
}
