// https://leetcode.com/problems/merge-two-sorted-lists/

package merge_two_sorted_lists

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2->2->4->5, 1->3->4->6
// 1->2->2->3->4->4->5->6
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	for l1 != nil {
		for l2 != nil {
			if l1.Val >= l2.Val {
				temp := &ListNode{
					Val:  l1.Val,
					Next: l1.Next,
				}
				l1.Val = l2.Val
				l1.Next = temp
				l2 = l2.Next
				break
			}
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	return l1
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
