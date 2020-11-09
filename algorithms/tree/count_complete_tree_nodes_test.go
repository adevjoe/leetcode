package tree

import "testing"

func TestCountNodes(t *testing.T) {
	cases := []struct {
		arg  *TreeNode
		desc string
		want int
	}{
		{
			arg:  &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 6}}},
			desc: "#1",
			want: 6,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := countNodes(s.arg); got != s.want {
				t.Errorf("want %d, get %d", s.want, got)
			}
		})
	}
}
