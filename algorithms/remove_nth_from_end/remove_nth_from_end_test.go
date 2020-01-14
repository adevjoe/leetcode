package remove_nth_from_end

import "testing"

func TestA(t *testing.T) {
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
	printListNode(link)
	link = removeNthFromEnd(link, 5)
	printListNode(link)
}
