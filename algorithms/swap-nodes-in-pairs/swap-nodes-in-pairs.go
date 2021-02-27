// https://leetcode.com/problems/swap-nodes-in-pairs

package leetcode

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := head.Next
	head.Next = swapPairs(n.Next)
	n.Next = head
	return n
}
