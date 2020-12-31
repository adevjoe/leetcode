// https://leetcode.com/problems/rotate-list/

package algorithms

import "github.com/adevjoe/leetcode/common"

func rotateRight(head *common.ListNode, k int) *common.ListNode {
	if k == 0 || head == nil {
		return head
	}
	rl, l := rotateList(head)
	var h1 *common.ListNode
	var h2 *common.ListNode
	rlHead := rl
	for k > l {
		k -= l
	}
	for i := 1; rl != nil && i <= k; i++ {
		if i == k {
			h2 = rl.Next
			rl.Next = nil
			h1 = rlHead
			break
		}
		rl = rl.Next
	}
	h1, _ = rotateList(h1)
	h2, _ = rotateList(h2)
	tail := h1
	for tail != nil {
		if tail.Next == nil {
			tail.Next = h2
			break
		}
		tail = tail.Next
	}
	return h1
}

func rotateList(head *common.ListNode) (*common.ListNode, int) {
	var h1 *common.ListNode
	i := 0
	for head != nil {
		i++
		newHead := &common.ListNode{}
		newHead.Val = head.Val
		newHead.Next = h1
		if head.Next == nil {
			return newHead, i
		}
		h1 = newHead
		head = head.Next
	}
	return nil, i
}

// alloc to many mem
func rotateRightIterator(head *common.ListNode, k int) *common.ListNode {
	if k == 0 || head == nil {
		return head
	}
	tmp := head
	for tmp != nil {
		if tmp.Next == nil {
			return head
		}
		if tmp.Next.Next == nil {
			tmp.Next.Next = head
			head = tmp.Next
			tmp.Next = nil
			break
		}
		tmp = tmp.Next
	}
	return rotateRightIterator(head, k-1)
}

func rotateRightCircle(head *common.ListNode, k int) *common.ListNode {
	tail := head
	l := 0
	for tail != nil {
		l++
		if tail.Next == nil {
			tail.Next = head
			break
		}
		tail = tail.Next
	}

	if k >= l {
		k = k % l
	}

	if k == 0 {
		tail.Next = nil
		return head
	}
	newHead := head
	for i := 1; i <= l-k; i++ {
		if i == l-k {
			t := newHead.Next
			newHead.Next = nil
			return t
		}
		newHead = newHead.Next
	}
	return nil
}
