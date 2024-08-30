package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	output := &ListNode{}
	currentOutputNode := output
	currentNode := head
	for currentNode.Next != nil {
		if currentNode.Val == 0 {
			if currentOutputNode.Val != 0 {
				newNode := &ListNode{}
				currentOutputNode.Next = newNode
				currentOutputNode = newNode
			}
		} else {
			currentOutputNode.Val += currentNode.Val
		}
		currentNode = currentNode.Next
	}
	return output
}

var _ = mergeNodes(nil)
