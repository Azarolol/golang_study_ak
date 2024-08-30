package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	len := 0
	currentNode := head
	for currentNode != nil {
		len++
		currentNode = currentNode.Next
	}
	values := make([]int, 0, len/2)
	currentNode = head
	for i := 0; i < len/2; i++ {
		values = append(values, currentNode.Val)
		currentNode = currentNode.Next
	}
	max := 0
	for i := 0; i < len/2; i++ {
		sum := values[len/2-i-1] + currentNode.Val
		if sum > max {
			max = sum
		}
		currentNode = currentNode.Next
	}
	return max
}

var _ = pairSum(nil)
