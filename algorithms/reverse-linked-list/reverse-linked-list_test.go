package leetcode

import (
	"reflect"
	"testing"
)

func TestReverseLinkedList(t *testing.T) {
	cases := []struct {
		desc string
		head *ListNode
		want *ListNode
	}{
		{
			desc: "#1",
		},
		{
			desc: "#2",
			head: &ListNode{Val: 1},
			want: &ListNode{Val: 1},
		},
		{
			desc: "#3",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			want: &ListNode{Val: 2, Next: &ListNode{Val: 1}},
		},
		{
			desc: "#4",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			want: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := recur(copyList(s.head)); !reflect.DeepEqual(s.want, got) {
				t.Error()
			}
			if got := iterate(s.head); !reflect.DeepEqual(s.want, got) {
				t.Error()
			}
		})
	}
}

func copyList(src *ListNode) *ListNode {
	if src == nil {
		return nil
	}
	loop := src
	for loop != nil {
		next := loop.Next
		loop.Next = &ListNode{
			Val:  loop.Val,
			Next: loop.Next,
		}
		loop = next
	}
	loop = src
	des := src.Next
	for loop != nil {
		if loop.Next.Next == nil {
			loop.Next = nil
			break
		}
		tmp := loop.Next
		loop.Next = loop.Next.Next
		loop = tmp
	}
	return des
}
