// https://leetcode.com/problems/add-two-numbers/

package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a := l1.Val + l2.Val
	if a >= 10 {
		l1.Val = a - 10
		if l2.Next == nil {
			l2.Next = new(ListNode)
		}
		// 有隐患， l2 的值会大于9
		l2.Next.Val += 1
	} else {
		l1.Val = a
	}
	if l2.Next != nil {
		if l1.Next == nil {
			l1.Next = new(ListNode)
		}
		l1.Next = AddTwoNumbers(l1.Next, l2.Next)
	}
	return l1
}
