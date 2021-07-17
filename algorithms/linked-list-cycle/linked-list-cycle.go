// https://leetcode.com/problems/linked-list-cycle

package leetcode

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	for head != nil {
		head.Val += 200001
		if head.Next != nil && head.Next.Val > 100000 {
			return true
		}
		head = head.Next
	}
	return false
}

func hasCycleWithTwoPoint(head *ListNode) bool {
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func hasCycleWithMap(head *ListNode) bool {
	m := make(map[*ListNode]bool, 0)
	for head != nil {
		m[head] = true
		if head.Next != nil && m[head.Next] {
			return true
		}
		head = head.Next
	}
	return false
}
