package leetcode

import (
	"reflect"
	"testing"
)

func TestIntersectionofTwoLinkedLists(t *testing.T) {
	cases := []struct {
		desc  string
		headA *ListNode
		headB *ListNode
		want  *ListNode
	}{
		{
			desc: "#1",
			headA: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 8,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			headB: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
		{
			desc: "#2",
			headA: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 4,
							},
						},
					},
				},
			},
			headB: &ListNode{
				Val: 3,
			},
		},
		{
			desc: "#3",
			headA: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 4,
					},
				},
			},
			headB: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	}

	cases[0].headB.Next.Next.Next = cases[0].headA.Next.Next
	cases[0].want = cases[0].headA.Next.Next
	cases[1].headB.Next = cases[1].headA.Next.Next.Next
	cases[1].want = cases[1].headA.Next.Next.Next

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := getIntersectionNode(s.headA, s.headB); !reflect.DeepEqual(got, s.want) {
				t.Error()
			}
		})
	}
}
