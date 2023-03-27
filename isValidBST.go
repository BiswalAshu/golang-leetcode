package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return checkValid(root, nil, nil)
}

func checkValid(root *TreeNode, min *int, max *int) bool {
	if root == nil {
		return true
	}
	if min != nil && *min >= root.Val {
		return false
	}
	if max != nil && *max <= root.Val {
		return false
	}
	return checkValid(root.Left, min, &root.Val) && checkValid(root.Right, &root.Val, max)
}
