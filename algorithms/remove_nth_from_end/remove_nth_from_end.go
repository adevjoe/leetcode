// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package remove_nth_from_end

import "fmt"

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// firstly, get list length
	l := 0
	var newHead *ListNode
	newHead = head
	for newHead != nil {
		newHead = newHead.Next
		l++
	}

	// get nth from start
	n = l - n + 1
	newHead = head
	// mark lastHead
	lastHead := head
	for i := 1; i <= n; i++ {
		if n == 1 {
			return newHead.Next
		}
		if i == n {
			// change lastHead next point
			lastHead.Next = newHead.Next
			break
		}
		lastHead = newHead
		newHead = newHead.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printListNode(head *ListNode) {
	var s string
	s += fmt.Sprint(head.Val)
	if head.Next == nil {
		fmt.Println(s)
		return
	}
	head = head.Next
	for head != nil {
		s += fmt.Sprintf("->%d", head.Val)
		head = head.Next
	}
	fmt.Println(s)
}
