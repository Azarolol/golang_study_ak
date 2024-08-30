package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	val, _, _ := walk(root, 0, 0, 0)
	return val
}

func walk(node *TreeNode, maxSum int, currentDepth int, maxDepth int) (int, int, int) {
	if node.Left == nil && node.Right == nil {
		if currentDepth > maxDepth {
			return node.Val, currentDepth, currentDepth
		} else if currentDepth == maxDepth {
			return node.Val + maxSum, currentDepth, currentDepth
		} else {
			return maxSum, currentDepth, maxDepth
		}
	} else if node.Left == nil {
		return walk(node.Right, maxSum, currentDepth+1, maxDepth)
	} else if node.Right == nil {
		return walk(node.Left, maxSum, currentDepth+1, maxDepth)
	} else {
		leftSum, leftCurrentDepth, leftMaxDepth := walk(node.Right, maxSum, currentDepth+1, maxDepth)
		rightSum, rightCurrentDepth, rightMaxDepth := walk(node.Left, maxSum, currentDepth+1, maxDepth)
		if leftMaxDepth > rightMaxDepth {
			return leftSum, leftCurrentDepth, leftMaxDepth
		} else if leftMaxDepth == rightMaxDepth {
			return leftSum + rightSum, leftCurrentDepth, leftMaxDepth
		} else {
			return rightSum, rightCurrentDepth, rightMaxDepth
		}
	}
}

var _ = deepestLeavesSum(nil)
