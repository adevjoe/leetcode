package merge_two_sorted_lists

import (
	"github.com/adevjoe/leetcode/common"
	"testing"
)

func TestA(t *testing.T) {
	l1 := &common.ListNode{Val: 2, Next: &common.ListNode{
		Val: 2,
		Next: &common.ListNode{
			Val: 4,
			Next: &common.ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}}
	l2 := &common.ListNode{Val: 1, Next: &common.ListNode{
		Val: 3,
		Next: &common.ListNode{
			Val: 4,
			Next: &common.ListNode{
				Val:  6,
				Next: nil,
			},
		},
	}}
	common.PrintListNode(l1)
	common.PrintListNode(l2)
	l3 := mergeTwoLists(l1, l2)
	common.PrintListNode(l3)
}
