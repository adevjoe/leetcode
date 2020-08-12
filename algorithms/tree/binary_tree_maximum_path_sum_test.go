package tree

import (
	"fmt"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	cases := []struct {
		arg  *TreeNode
		want int
	}{
		{arg: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, want: 6},
		{arg: &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, want: 42},
		{arg: &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: -8, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, want: 15},
		{arg: &TreeNode{Val: -10, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: -11}, Right: &TreeNode{Val: 7}}}, want: 27},
		{arg: &TreeNode{Val: -8, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: -11}, Right: &TreeNode{Val: 7}}}, want: 28},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("TestMaxPathSum-%d", i), func(t *testing.T) {
			if get := maxPathSum(c.arg); get != c.want {
				t.Errorf("arg: %v, get: %v, want: %v", preorderTraversal(c.arg), get, c.want)
			}
		})
	}
}
