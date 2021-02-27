package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotateList(t *testing.T) {
	cases := []struct {
		desc string
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			desc: "#1",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			k:    2,
			want: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{
			desc: "#2",
			head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			k:    4,
			want: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
		{
			desc: "#3",
			head: &ListNode{Val: 1},
			k:    99,
			want: &ListNode{Val: 1},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := rotateRight(s.head, s.k); !reflect.DeepEqual(s.want, got) {
				t.Errorf("")
				PrintListNode(s.head)
				PrintListNode(s.want)
				PrintListNode(got)
			}
		})
	}
}
func TestRotateRightIterator(t *testing.T) {
	cases := []struct {
		desc string
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			desc: "#1",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			k:    2,
			want: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{
			desc: "#2",
			head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			k:    4,
			want: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
		{
			desc: "#3",
			head: &ListNode{Val: 1},
			k:    99,
			want: &ListNode{Val: 1},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := rotateRightIterator(s.head, s.k); !reflect.DeepEqual(s.want, got) {
				t.Errorf("")
				PrintListNode(s.head)
				PrintListNode(s.want)
				PrintListNode(got)
			}
		})
	}
}
func TestRotateCircle(t *testing.T) {
	cases := []struct {
		desc string
		head *ListNode
		k    int
		want *ListNode
	}{
		{
			desc: "#1",
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			k:    2,
			want: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}},
		},
		{
			desc: "#2",
			head: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
			k:    4,
			want: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}},
		},
		{
			desc: "#3",
			head: &ListNode{Val: 1},
			k:    99,
			want: &ListNode{Val: 1},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := rotateRightCircle(s.head, s.k); !reflect.DeepEqual(s.want, got) {
				t.Errorf("")
				PrintListNode(s.head)
				PrintListNode(s.want)
				PrintListNode(got)
			}
		})
	}
}

func BenchmarkRotateRightCircle(b *testing.B) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	for i := 0; i < b.N; i++ {
		rotateRightCircle(head, 2)
	}
}
func BenchmarkRotateRight(b *testing.B) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	for i := 0; i < b.N; i++ {
		rotateRight(head, 2)
	}
}
func BenchmarkRotateRightIterator(b *testing.B) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	for i := 0; i < b.N; i++ {
		rotateRightIterator(head, 2)
	}
}
func BenchmarkRotateRightCircle2000(b *testing.B) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	for i := 0; i < b.N; i++ {
		rotateRightCircle(head, 2000)
	}
}
func BenchmarkRotateRight2000(b *testing.B) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	for i := 0; i < b.N; i++ {
		rotateRight(head, 2000)
	}
}
func BenchmarkRotateRightIterator2000(b *testing.B) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	for i := 0; i < b.N; i++ {
		rotateRightIterator(head, 2000)
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
