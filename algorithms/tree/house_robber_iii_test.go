// +build !ci

package tree

import "testing"

func TestRob(t *testing.T) {
	cases := []struct {
		desc string
		root *TreeNode
		want int
	}{
		{
			desc: "#1",
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}},
			want: 7,
		},
		{
			desc: "#2",
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 1}}},
			want: 9,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := rob(s.root); got != s.want {
				t.Errorf("root: %v, want: %d, got: %d", inorderTraversal(s.root), s.want, got)
			}
		})
	}
}
