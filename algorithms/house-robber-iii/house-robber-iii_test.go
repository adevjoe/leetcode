package leetcode

import "testing"

func TestHouseRobberIII(t *testing.T) {
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
				t.Errorf("want: %d, got: %d", s.want, got)
			}
		})

		t.Run(s.desc, func(t *testing.T) {
			if got := robTwo(s.root); got != s.want {
				t.Errorf("want: %d, got: %d", s.want, got)
			}
		})
	}
}
