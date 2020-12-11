// +build !ci

package algorithms

import (
	"fmt"
	"github.com/adevjoe/leetcode/common"
	"testing"
)

func TestPartition(t *testing.T) {
	cases := []struct {
		desc string
		head *common.ListNode
		x    int
		want *common.ListNode
	}{
		{
			desc: "#1",
			head: &common.ListNode{
				Val: 1, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 5, Next: &common.ListNode{Val: 2}}}}},
			},
			x: 3,
			want: &common.ListNode{
				Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 5}}}}},
			},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := partition(s.head, s.x); !common.IsSameListNode(got, s.want) {
				fmt.Println("head")
				common.PrintListNode(s.head)
				fmt.Println("want")
				common.PrintListNode(s.want)
				fmt.Println("got")
				common.PrintListNode(got)
				t.Errorf("x: %d", s.x)
			}
		})
	}
}
