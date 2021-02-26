package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	result := addTwoNumbers(l1, l2)
	printListNode(result)
}

func printListNode(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Print(head.Val)
	if head.Next != nil {
		fmt.Print(" -> ")
		printListNode(head.Next)
	} else {
		fmt.Print("\n")
	}
}
