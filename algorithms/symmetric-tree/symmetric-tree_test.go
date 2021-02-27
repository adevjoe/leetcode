package leetcode

import (
	"testing"
)

func TestSymmetricTree(t *testing.T) {
	cases := []struct {
		desc string
		root *TreeNode
		want bool
	}{
		{
			desc: "#1",
			root: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}},
			want: true,
		},
		{
			desc: "#2",
			root: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
			want: false,
		},
		{
			desc: "#3",
			root: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3}},
			want: false,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := isSymmetric(s.root); got != s.want {
				t.Error()
			}
		})
	}
}
