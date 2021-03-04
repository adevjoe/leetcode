package leetcode

import (
	"testing"
)

func TestMinimumDepthofBinaryTree(t *testing.T) {
	cases := []struct {
		desc string
		args *TreeNode
		want int
	}{
		{desc: "#1", args: &TreeNode{Val: 1}, want: 1},
		{
			desc: "#2",
			args: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			}},
			want: 2,
		},
		{desc: "#3", args: &TreeNode{Val: 1}, want: 1},
		{desc: "#4", args: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, want: 2},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := minDepth(s.args); s.want != got {
				t.Error()
			}
		})
	}
}
