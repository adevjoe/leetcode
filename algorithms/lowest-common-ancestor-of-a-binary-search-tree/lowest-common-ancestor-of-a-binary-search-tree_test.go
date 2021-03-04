package leetcode

import (
	"reflect"
	"testing"
)

func TestLowestCommonAncestorofaBinarySearchTree(t *testing.T) {
	cases := []struct {
		desc string
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			desc: "#1",
			root: &TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}},
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 8},
			want: &TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}},
		},
		{
			desc: "#2",
			root: &TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}},
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 4},
			want: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}},
		},
		{
			desc: "#3",
			root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
			p:    &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
			q:    &TreeNode{Val: 1},
			want: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
		},
		{
			desc: "#4",
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}},
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 3},
			want: &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := lowestCommonAncestor(s.root, s.p, s.q); !reflect.DeepEqual(got, s.want) {
				t.Error()
			}
		})
		t.Run(s.desc, func(t *testing.T) {
			if got := lowestCommonAncestor2(s.root, s.p, s.q); !reflect.DeepEqual(got, s.want) {
				t.Error()
			}
		})
	}
}
