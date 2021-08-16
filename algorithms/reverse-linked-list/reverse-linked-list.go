// https://leetcode.com/problems/reverse-linked-list

package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func recur(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	r := head
	for head.Next != nil {
		r = recur(head.Next)
		head.Next.Next = head
		head.Next = nil
	}
	return r
}

func iterate(head *ListNode) *ListNode {
	var tail *ListNode
	for head != nil {
		next := head.Next
		head.Next = tail
		tail = head
		head = next
	}
	return tail
}
