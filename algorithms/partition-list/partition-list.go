// https://leetcode.com/problems/partition-list

package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var l1Head = &ListNode{}
	var l2Head = &ListNode{}
	l1 := l1Head
	l2 := l2Head
	for head != nil {
		if head.Val < x {
			l1.Next = head
			l1 = l1.Next
		} else {
			l2.Next = head
			l2 = l2.Next
		}
		head = head.Next
	}
	l2.Next = nil
	l1.Next = l2Head.Next
	return l1Head.Next
}
