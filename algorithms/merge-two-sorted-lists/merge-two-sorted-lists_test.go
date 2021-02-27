package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeTwoSortedLists(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}}
	l2 := &ListNode{Val: 1, Next: &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}}
	PrintListNode(l1)
	PrintListNode(l2)
	l3 := mergeTwoLists(l1, l2)
	PrintListNode(l3)
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
