package tree

import "testing"

func TestKthSmallest(t *testing.T) {
	cases := []struct {
		desc string
		arg1 *TreeNode
		arg2 int
		want int
	}{
		{
			desc: "#1",
			arg1: &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}},
			arg2: 1,
			want: 1,
		},
		{
			desc: "#2",
			arg1: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6}},
			arg2: 3,
			want: 3,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := kthSmallest(s.arg1, s.arg2); s.want != got {
				t.Errorf("arg1: %v, arg2: %v, got: %v, want: %v", inorderTraversal(s.arg1), s.arg2, got, s.want)
			}
		})
	}
}
