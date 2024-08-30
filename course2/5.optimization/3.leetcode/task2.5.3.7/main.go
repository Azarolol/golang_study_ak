package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	arr := collectValues(root)
	return makeTree(arr)
}

func collectValues(node *TreeNode) []int {
	arr := make([]int, 0)
	if node.Left != nil {
		arr = append(arr, collectValues(node.Left)...)
	}
	arr = append(arr, node.Val)
	if node.Right != nil {
		arr = append(arr, collectValues(node.Right)...)
	}
	return arr
}

func makeTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	if len(arr) == 1 {
		return &TreeNode{Val: arr[0]}
	}
	mid := len(arr) / 2
	return &TreeNode{Val: arr[mid], Left: makeTree(arr[:mid]), Right: makeTree(arr[mid+1:])}
}

var _ = balanceBST(nil)
