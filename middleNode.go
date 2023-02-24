package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// make slice of list nodes, return slice (len by 2)
func middleNode(head *ListNode) *ListNode {
	nodesSlice := make([]*ListNode, 0)
	for head != nil {
		nodesSlice = append(nodesSlice, head)
		head = head.Next
	}
	return nodesSlice[len(nodesSlice)/2]
}
