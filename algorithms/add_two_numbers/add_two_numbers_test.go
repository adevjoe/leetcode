package add_two_numbers

import (
	"github.com/adevjoe/leetcode/common"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &common.ListNode{Val: 2, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 3}}}
	l2 := &common.ListNode{Val: 5, Next: &common.ListNode{Val: 6, Next: &common.ListNode{Val: 4}}}
	common.PrintListNode(l1)
	common.PrintListNode(l2)
	result := addTwoNumbers(l1, l2)
	common.PrintListNode(result)
}
