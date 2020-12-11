package common

import "fmt"

// 单链表 singly-linked list
type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func PrintListNode(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Print(head.Val)
	if head.Next != nil {
		fmt.Print(" -> ")
		PrintListNode(head.Next)
	} else {
		fmt.Print("\n")
	}
}

func IsSameListNode(l, n *ListNode) bool {
	for l != nil && n != nil {
		if l.Val != n.Val {
			return false
		} else {
			l = l.Next
			n = n.Next
			continue
		}
	}
	if l == nil && n == nil {
		return true
	}
	return false
}
