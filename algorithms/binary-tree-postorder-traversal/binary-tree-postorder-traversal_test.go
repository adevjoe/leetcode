package leetcode

import (
	"reflect"
	"testing"
)

func TestBinaryTreePostorderTraversal(t *testing.T) {
	cases := []struct {
		desc string
		root *TreeNode
		want []int
	}{
		{
			desc: "#1",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:   6,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			want: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			desc: "#2",
			want: nil,
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := postorderTraversal(s.root); !reflect.DeepEqual(s.want, got) {
				t.Error()
			}
		})
	}
}
