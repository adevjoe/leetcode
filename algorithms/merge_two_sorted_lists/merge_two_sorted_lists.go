// https://leetcode.com/problems/merge-two-sorted-lists/

package merge_two_sorted_lists

import (
	"github.com/adevjoe/leetcode/common"
)

// 2->2->4->5, 1->3->4->6
// 1->2->2->3->4->4->5->6
func mergeTwoLists(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	var newList *common.ListNode
	for l1 != nil || l2 != nil {
		tail := &common.ListNode{}
		if l1 != nil {
			if l2 != nil {
				if l1.Val.(int) <= l2.Val.(int) {
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
