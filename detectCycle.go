package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	cycleExists := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			cycleExists = true
			break
		}
	}
	if !cycleExists {
		return nil
	}
	for head != slow {
		head = head.Next
		slow = slow.Next
	}
	return head
}
