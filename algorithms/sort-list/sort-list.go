// https://leetcode.com/problems/sort-list

package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	done := head == nil
	list := make([]*ListNode, 2)

	for step := 1; !done; step *= 2 {
		done = true
		prev := dummy
		remaining := prev.Next
		for remaining != nil {
			for i := 0; i < 2; i++ {
				list[i] = remaining
				var tail *ListNode
				for j := 0; j < step && remaining != nil; j++ {
					tail = remaining
					remaining = remaining.Next
				}
				if tail != nil {
					tail.Next = nil
				}
			}

			done = done && remaining == nil

			if list[1] != nil {
				for list[0] != nil || list[1] != nil {
					idx := 1
					if nil == list[1] || nil != list[0] && list[0].Val <= list[1].Val {
						idx = 0
					}
					prev.Next = list[idx]
					list[idx] = list[idx].Next
					prev = prev.Next
				}

				prev.Next = nil
			} else {
				prev.Next = list[0]
			}
		}
	}
	return dummy.Next
}

func sortListBubble(head *ListNode) *ListNode {
	loop := head
	for loop != nil {
		cur := loop.Next
		for cur != nil {
			if cur.Val < loop.Val {
				cur.Val, loop.Val = loop.Val, cur.Val
			}
			cur = cur.Next
		}
		loop = loop.Next
	}
	return head
}
