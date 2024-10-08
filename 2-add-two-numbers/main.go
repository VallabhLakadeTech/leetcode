package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 999
// 999
// 8 8 8 1

//243
//564

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	firstNode := ListNode{}
	currentNode := &firstNode
	prevNode := firstNode
	carryForward := 0
	for l1 != nil || l2 != nil || carryForward != 0 {
		if currentNode == nil {
			currentNode = &ListNode{Val: carryForward}
			prevNode.Next = currentNode

		}
		if l1 == nil && l2 == nil {
			currentNode.Val = carryForward
			carryForward = 0
		}
		if l1 != nil {
			currentNode.Val = currentNode.Val + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			currentNode.Val = currentNode.Val + l2.Val
			l2 = l2.Next
		}
		if currentNode.Val >= 10 {
			currentNode.Val = currentNode.Val % 10
			carryForward = carryForward + 1
		} else {
			carryForward = 0
		}
		currentNode = currentNode.Next
	}
	fmt.Println(prevNode)
	return &firstNode
}
