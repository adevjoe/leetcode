package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCopyListwithRandomPointer(t *testing.T) {
	cases := []struct {
		desc string
		head *Node
		want *Node
	}{
		{
			desc: "#1",
			head: &Node{
				Val: 7,
				Next: &Node{
					Val: 13,
					Next: &Node{
						Val: 11,
						Next: &Node{
							Val: 10,
							Next: &Node{
								Val: 1,
							},
						},
					},
				},
			},
			want: &Node{
				Val: 7,
				Next: &Node{
					Val: 13,
					Next: &Node{
						Val: 11,
						Next: &Node{
							Val: 10,
							Next: &Node{
								Val: 1,
							},
						},
					},
				},
			},
		},
		{
			desc: "#2",
			head: &Node{
				Val: 1,
				Next: &Node{
					Val: 2,
				},
			},
			want: &Node{
				Val: 1,
				Next: &Node{
					Val: 2,
				},
			},
		},
		{
			desc: "#3",
			head: &Node{
				Val: 3,
				Next: &Node{
					Val: 3,
					Next: &Node{
						Val: 3,
					},
				},
			},
			want: &Node{
				Val: 3,
				Next: &Node{
					Val: 3,
					Next: &Node{
						Val: 3,
					},
				},
			},
		},
		{
			desc: "#4",
		},
	}
	cases[0].head.Next.Random = cases[0].head
	cases[0].head.Next.Next.Random = cases[0].head.Next.Next.Next.Next
	cases[0].head.Next.Next.Next.Next.Random = cases[0].head
	cases[0].head.Next.Next.Next.Random = cases[0].head.Next.Next
	cases[0].want.Next.Random = cases[0].want
	cases[0].want.Next.Next.Random = cases[0].want.Next.Next.Next.Next
	cases[0].want.Next.Next.Next.Next.Random = cases[0].want
	cases[0].want.Next.Next.Next.Random = cases[0].want.Next.Next

	cases[1].head.Random = cases[1].head.Next
	cases[1].head.Next.Random = cases[1].head.Next
	cases[1].want.Random = cases[1].want.Next
	cases[1].want.Next.Random = cases[1].want.Next

	cases[2].head.Next.Random = cases[2].head
	cases[2].want.Next.Random = cases[2].want
	for _, s := range cases {
		if got := printNode(copyRandomList(s.head)); got != printNode(s.want) {
			t.Errorf("want: %s, got: %s", printNode(s.want), got)
		}
		if got := printNode(copyRandomList2(s.head)); got != printNode(s.want) {
			t.Errorf("want: %s, got: %s", printNode(s.want), got)
		}
	}
}

func printNode(head *Node) string {
	if head == nil {
		return ""
	}
	result := strconv.Itoa(head.Val)
	if head.Random != nil {
		result += fmt.Sprintf("(%d)", head.Random.Val)
	}
	if head.Next != nil {
		result += " -> "
		result += printNode(head.Next)
	}
	return result
}
