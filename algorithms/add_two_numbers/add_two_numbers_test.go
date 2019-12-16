package add_two_numbers

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	printListNode(l1)
	printListNode(l2)
	result := addTwoNumbers(l1, l2)
	printListNode(result)
}

func printListNode(node *ListNode) {
	fmt.Print(node.Val)
	if node.Next != nil {
		fmt.Print(" -> ")
		printListNode(node.Next)
	} else {
		fmt.Print("\n")
	}
}
