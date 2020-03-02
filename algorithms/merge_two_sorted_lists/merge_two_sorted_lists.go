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
	var newList *ListNode
	for l1 != nil || l2 != nil {
		tail := &ListNode{}
		if l1 != nil {
			if l2 != nil {
				if l1.Val <= l2.Val {
					tail.Val = l1.Val
					l1 = l1.Next
				} else {
					tail.Val = l2.Val
					l2 = l2.Next
				}
			} else {
				tail.Val = l1.Val
				l1 = l1.Next
			}
		} else {
			if l2 != nil {
				tail.Val = l2.Val
				l2 = l2.Next
			}
		}
		if newList == nil {
			newList = tail
		} else {
			tempList := newList
			last := tempList
			for tempList != nil {
				last = tempList
				tempList = tempList.Next
			}
			last.Next = tail
		}
	}
	return newList
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
