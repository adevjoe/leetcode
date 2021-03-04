package leetcode

import (
	"testing"
)

func TestBalancedBinaryTree(t *testing.T) {
	cases := []struct {
		desc string
		root *TreeNode
		want bool
	}{
		{
			desc: "#1",
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			want: true,
		},
		{
			desc: "#2",
			root: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 2},
			},
			want: false,
		},
		{
			desc: "#3",
			root: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}},
				Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
			},
			want: false,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := isBalanced(s.root); got != s.want {
				t.Errorf("got: %v, want: %v", got, s.want)
			}
		})
	}
}
