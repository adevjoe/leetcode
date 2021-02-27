package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartitionList(t *testing.T) {
	cases := []struct {
		desc string
		head *ListNode
		x    int
		want *ListNode
	}{
		{
			desc: "#1",
			head: &ListNode{
				Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}},
			},
			x: 3,
			want: &ListNode{
				Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}},
			},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := partition(s.head, s.x); !reflect.DeepEqual(got, s.want) {
				fmt.Println("head")
				PrintListNode(s.head)
				fmt.Println("want")
				PrintListNode(s.want)
				fmt.Println("got")
				PrintListNode(got)
				t.Errorf("x: %d", s.x)
			}
		})
	}
}

func PrintListNode(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Print(head.Val)
	if head.Next != nil {
		fmt.Print(" -> ")
		PrintListNode(head.Next)
	} else {
		fmt.Print("\n")
	}
}
