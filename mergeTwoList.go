package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	mergedList := &ListNode{}
	returnList := mergedList

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			// node := &ListNode{Val: list1.Val, Next: nil}
			mergedList.Next = list1
			list1 = list1.Next
		} else {
			// node := &ListNode{Val: list2.Val, Next: nil}
			mergedList.Next = list2
			list2 = list2.Next
		}
		mergedList = mergedList.Next
	}
	if list1 != nil {
		mergedList.Next = list1
	}
	if list2 != nil {
		mergedList.Next = list2
	}

	return returnList.Next
}
