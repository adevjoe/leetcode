package leetcode

import (
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	cases := []struct {
		desc string
		head *ListNode
		want *ListNode
	}{
		{
			desc: "#1",
			head: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 1,
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
		},
		{
			desc: "#2",
			head: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 0,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: -1,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			desc: "#3",
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := sortListBubble(s.head); !reflect.DeepEqual(s.want, got) {
				t.Error()
			}
		})
	}
	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := sortList(s.head); !reflect.DeepEqual(s.want, got) {
				t.Error()
			}
		})
	}
}
