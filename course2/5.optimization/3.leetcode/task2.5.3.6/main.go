package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	maxValue := nums[0]
	maxI := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			maxI = i
		}
	}
	newNode := &TreeNode{Val: maxValue}
	if maxI != 0 {
		left := constructMaximumBinaryTree(nums[:maxI])
		newNode.Left = left
	}
	if maxI != len(nums)-1 {
		right := constructMaximumBinaryTree(nums[maxI+1:])
		newNode.Right = right
	}
	return newNode
}

var _ = constructMaximumBinaryTree(nil)
