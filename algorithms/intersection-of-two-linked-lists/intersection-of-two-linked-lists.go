// https://leetcode.com/problems/intersection-of-two-linked-lists

package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	loopA := headA
	loopB := headB
	lenA := 0
	lenB := 0
	for loopA != nil || loopB != nil {
		if loopA == loopB {
			return loopA
		}
		if loopA != nil {
			lenA++
			loopA = loopA.Next
		}
		if loopB != nil {
			lenB++
			loopB = loopB.Next
		}
	}

	loopA = headA
	loopB = headB
	if lenA > lenB {
		for i := 1; i <= lenA-lenB; i++ {
			loopA = loopA.Next
		}
	} else {
		for i := 1; i <= lenB-lenA; i++ {
			loopB = loopB.Next
		}
	}
	for loopA != nil {
		if loopA == loopB {
			return loopA
		}
		loopA = loopA.Next
		loopB = loopB.Next
	}
	return nil
}
