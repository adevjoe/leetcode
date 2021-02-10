package merge_k_sorted_lists

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	var list []*ListNode
	list = append(list, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	})
	list = append(list, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	})
	list = append(list, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	})

	cases := []struct {
		desc  string
		lists []*ListNode
		want  *ListNode
	}{
		{
			desc:  "#1",
			lists: list,
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}},
				},
			},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := mergeKLists(s.lists); !sameListNode(s.want, got) {
				t.Error("error")
				fmt.Print("want: ")
				printListNode(s.want)
				fmt.Print("got: ")
				printListNode(got)
			}
		})
	}
}

func printListNode(head *ListNode) {
	fmt.Print(head.Val)
	if head.Next != nil {
		fmt.Print(" -> ")
		printListNode(head.Next)
	} else {
		fmt.Print("\n")
	}
}

func sameListNode(l1 *ListNode, l2 *ListNode) bool {
	if l1 == nil && l1 != l2 {
		return false
	}
	for l1 != nil {
		if l2 == nil {
			return false
		}
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return true
}
