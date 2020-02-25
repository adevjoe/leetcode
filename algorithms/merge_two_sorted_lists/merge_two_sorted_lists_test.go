package merge_two_sorted_lists

import "testing"

func TestA(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}}
	l2 := &ListNode{Val: 1, Next: &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}}
	printListNode(l1)
	printListNode(l2)
	l3 := mergeTwoLists(l1, l2)
	printListNode(l3)
}
