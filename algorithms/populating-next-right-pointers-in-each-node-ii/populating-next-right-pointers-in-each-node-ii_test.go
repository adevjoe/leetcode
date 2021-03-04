package leetcode

import (
	"reflect"
	"testing"
)

func TestPopulatingNextRightPointersinEachNodeII(t *testing.T) {
	cases := []struct {
		desc string
		root *Node
		want *Node
	}{
		{
			desc: "#1",
			root: &Node{Val: 1, Left: &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}}, Right: &Node{Val: 3, Right: &Node{Val: 7}}},
			want: &Node{Val: 1, Left: &Node{Val: 2, Left: &Node{Val: 4}, Right: &Node{Val: 5}}, Right: &Node{Val: 3, Right: &Node{Val: 7}}},
		},
	}
	cases[0].want.Left.Next = cases[0].want.Right
	cases[0].want.Left.Left.Next = cases[0].want.Left.Right
	cases[0].want.Left.Right.Next = cases[0].want.Right.Right

	for _, s := range cases {
		if got := connect(s.root); !reflect.DeepEqual(got, s.want) {
			t.Error()
		}
	}
}
