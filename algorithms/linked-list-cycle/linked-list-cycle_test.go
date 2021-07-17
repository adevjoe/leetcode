package leetcode

import "testing"

func TestLinkedListCycle(t *testing.T) {
	cases := []struct {
		desc string
		head *ListNode
		want bool
	}{
		{
			desc: "#1",
			head: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: -4,
						},
					},
				},
			},
			want: true,
		},
		{
			desc: "#2",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			want: true,
		},
		{
			desc: "#3",
			head: &ListNode{
				Val: 1,
			},
			want: false,
		},
	}

	cases[0].head.Next.Next.Next = cases[0].head.Next
	cases[1].head.Next = cases[1].head

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := hasCycle(s.head); got != s.want {
				t.Error()
			}
			if got := hasCycleWithTwoPoint(s.head); got != s.want {
				t.Error()
			}
			if got := hasCycleWithMap(s.head); got != s.want {
				t.Error()
			}
		})
	}
}
