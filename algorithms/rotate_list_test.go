package algorithms

import (
	"github.com/adevjoe/leetcode/common"
	"testing"
)

func TestRotateRight(t *testing.T) {
	cases := []struct {
		desc string
		head *common.ListNode
		k    int
		want *common.ListNode
	}{
		{
			desc: "#1",
			head: &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 5}}}}},
			k:    2,
			want: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 5, Next: &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3}}}}},
		},
		{
			desc: "#2",
			head: &common.ListNode{Val: 0, Next: &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2}}},
			k:    4,
			want: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 0, Next: &common.ListNode{Val: 1}}},
		},
		{
			desc: "#3",
			head: &common.ListNode{Val: 1},
			k:    99,
			want: &common.ListNode{Val: 1},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := rotateRight(s.head, s.k); !common.IsSameListNode(s.want, got) {
				t.Errorf("")
				common.PrintListNode(s.head)
				common.PrintListNode(s.want)
				common.PrintListNode(got)
			}
		})
	}
}
func TestRotateRightIterator(t *testing.T) {
	cases := []struct {
		desc string
		head *common.ListNode
		k    int
		want *common.ListNode
	}{
		{
			desc: "#1",
			head: &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 5}}}}},
			k:    2,
			want: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 5, Next: &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3}}}}},
		},
		{
			desc: "#2",
			head: &common.ListNode{Val: 0, Next: &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2}}},
			k:    4,
			want: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 0, Next: &common.ListNode{Val: 1}}},
		},
		{
			desc: "#3",
			head: &common.ListNode{Val: 1},
			k:    99,
			want: &common.ListNode{Val: 1},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := rotateRightIterator(s.head, s.k); !common.IsSameListNode(s.want, got) {
				t.Errorf("")
				common.PrintListNode(s.head)
				common.PrintListNode(s.want)
				common.PrintListNode(got)
			}
		})
	}
}
func TestRotateCircle(t *testing.T) {
	cases := []struct {
		desc string
		head *common.ListNode
		k    int
		want *common.ListNode
	}{
		{
			desc: "#1",
			head: &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 5}}}}},
			k:    2,
			want: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 5, Next: &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3}}}}},
		},
		{
			desc: "#2",
			head: &common.ListNode{Val: 0, Next: &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2}}},
			k:    4,
			want: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 0, Next: &common.ListNode{Val: 1}}},
		},
		{
			desc: "#3",
			head: &common.ListNode{Val: 1},
			k:    99,
			want: &common.ListNode{Val: 1},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			got := rotateRightCircle(s.head, s.k)
			common.PrintListNode(s.want)
			common.PrintListNode(got)
		})
	}
}

func BenchmarkRotateRightCircle(b *testing.B) {
	head := &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 5}}}}}

	for i := 0; i < b.N; i++ {
		rotateRightCircle(head, 2)
	}
}
func BenchmarkRotateRight(b *testing.B) {
	head := &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 5}}}}}

	for i := 0; i < b.N; i++ {
		rotateRight(head, 2)
	}
}
func BenchmarkRotateRightIterator(b *testing.B) {
	head := &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 5}}}}}

	for i := 0; i < b.N; i++ {
		rotateRightIterator(head, 2)
	}
}
func BenchmarkRotateRightCircle2000(b *testing.B) {
	head := &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 5}}}}}

	for i := 0; i < b.N; i++ {
		rotateRightCircle(head, 2000)
	}
}
func BenchmarkRotateRight2000(b *testing.B) {
	head := &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 5}}}}}

	for i := 0; i < b.N; i++ {
		rotateRight(head, 2000)
	}
}
func BenchmarkRotateRightIterator2000(b *testing.B) {
	head := &common.ListNode{Val: 1, Next: &common.ListNode{Val: 2, Next: &common.ListNode{Val: 3, Next: &common.ListNode{Val: 4, Next: &common.ListNode{Val: 5}}}}}

	for i := 0; i < b.N; i++ {
		rotateRightIterator(head, 2000)
	}
}
