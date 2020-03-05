// https://leetcode.com/problems/remove-nth-node-from-end-of-list/

package remove_nth_from_end

import (
	"github.com/adevjoe/leetcode/common"
)

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */
func removeNthFromEnd(head *common.ListNode, n int) *common.ListNode {
	// firstly, get list length
	l := 0
	var newHead *common.ListNode
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
