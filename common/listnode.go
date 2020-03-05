package common

import "fmt"

// 单链表 singly-linked list
type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func PrintListNode(head *ListNode) {
	fmt.Print(head.Val)
	if head.Next != nil {
		fmt.Print(" -> ")
		PrintListNode(head.Next)
	} else {
		fmt.Print("\n")
	}
}
