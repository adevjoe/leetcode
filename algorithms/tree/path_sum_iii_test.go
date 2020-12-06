package tree

import "testing"

func TestPathSumIII(t *testing.T) {
	cases := []struct {
		desc string
		root *TreeNode
		sum  int
		want int
	}{
		{
			desc: "#1",
			root: &TreeNode{
				Val: 10, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: -2}}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 1}}}, Right: &TreeNode{Val: -3, Right: &TreeNode{Val: 11}},
			},
			sum:  8,
			want: 3,
		},
		{
			desc: "#2",
			root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}}}}},
			sum:  3,
			want: 2,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := pathSumIII(s.root, s.sum); got != s.want {
				t.Errorf("want: %d, got: %d", s.want, got)
			}
		})
	}
}
