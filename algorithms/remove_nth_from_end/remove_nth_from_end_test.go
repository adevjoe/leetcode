package remove_nth_from_end

import (
	"github.com/adevjoe/leetcode/common"
	"testing"
)

func TestA(t *testing.T) {
	link := &common.ListNode{Val: 1, Next: &common.ListNode{
		Val: 2,
		Next: &common.ListNode{
			Val: 3,
			Next: &common.ListNode{
				Val: 4,
				Next: &common.ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}}
	common.PrintListNode(link)
	link = removeNthFromEnd(link, 5)
	common.PrintListNode(link)
}
