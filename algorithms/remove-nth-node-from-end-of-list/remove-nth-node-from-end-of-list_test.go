package leetcode

import (
	"testing"
)

func TestRemoveNthNodeFromEndofList(t *testing.T) {
	link := &ListNode{Val: 1, Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}}
	link = removeNthFromEnd(link, 5)
}
