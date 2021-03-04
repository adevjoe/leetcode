package tree

import "testing"

func TestSumOfLeftLeaves(t *testing.T) {
	cases := []struct {
		desc string
		root *TreeNode
		want int
	}{
		{
			desc: "#1",
			root: &TreeNode{
				Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
			},
			want: 24,
		},
		{
			desc: "#2",
			root: &TreeNode{
				Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}},
			},
			want: 24,
		},
		{
			desc: "#3",
			root: &TreeNode{
				Val: 3,
			},
			want: 0,
		},
		{
			desc: "#4",
			root: &TreeNode{
				Val: 3, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}},
			},
			want: 19,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := sumOfLeftLeaves(s.root); got != s.want {
				t.Error()
			}
		})
	}
}
