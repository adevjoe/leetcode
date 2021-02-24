package merge_k_sorted_lists

import (
	"fmt"
	"reflect"
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
			if got := mergeKLists(s.lists); !reflect.DeepEqual(s.want, got) {
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
