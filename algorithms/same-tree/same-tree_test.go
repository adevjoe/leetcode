package leetcode

import "testing"

func TestSameTree(t *testing.T) {
	cases := []struct {
		desc string
		t1   *TreeNode
		t2   *TreeNode
		want bool
	}{
		{
			desc: "#1",
			t1:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			t2:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			want: true,
		},
		{
			desc: "#2",
			t1:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			t2:   &TreeNode{Val: 1, Right: &TreeNode{Val: 2}},
			want: false,
		},
		{
			desc: "#3",
			t1:   &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}},
			t2:   &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			want: false,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := isSameTree(s.t1, s.t2); got != s.want {
				t.Errorf("error")
			}
		})
	}
}
