// https://leetcode.com/problems/partition-list/

package algorithms

import "github.com/adevjoe/leetcode/common"

func partition(head *common.ListNode, x int) *common.ListNode {
	var l1Head = &common.ListNode{}
	var l2Head = &common.ListNode{}
	l1 := l1Head
	l2 := l2Head
	for head != nil {
		if head.Val.(int) < x {
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
