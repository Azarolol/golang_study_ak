package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	walk(root, 0)
	return root
}

func walk(node *TreeNode, sum int) int {
	if node.Right != nil {
		sum = walk(node.Right, sum)
	}
	sum += node.Val
	node.Val = sum
	if node.Left != nil {
		sum = walk(node.Left, sum)
	}
	return sum
}

var _ = bstToGst(nil)
