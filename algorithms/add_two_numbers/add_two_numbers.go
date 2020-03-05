// https://leetcode.com/problems/add-two-numbers/

package add_two_numbers

import "github.com/adevjoe/leetcode/common"

func addTwoNumbers(l1 *common.ListNode, l2 *common.ListNode) *common.ListNode {
	a := l1.Val.(int) + l2.Val.(int)
	if a >= 10 {
		l1.Val = a - 10
		if l2.Next == nil {
			l2.Next = new(common.ListNode)
		}
		// 有隐患， l2 的值会大于9
		l2.Next.Val = l2.Next.Val.(int) + 1
	} else {
		l1.Val = a
	}
	if l2.Next != nil {
		if l1.Next == nil {
			l1.Next = new(common.ListNode)
		}
		l1.Next = addTwoNumbers(l1.Next, l2.Next)
	}
	return l1
}
