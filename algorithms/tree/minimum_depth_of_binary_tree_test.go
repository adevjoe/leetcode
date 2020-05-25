package tree

import (
	"fmt"
	"testing"
)

func TestMinDepth(t *testing.T) {
	cases := []struct {
		args *TreeNode
		want int
	}{
		{args: &TreeNode{Val: 1}, want: 1},
		{
			args: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			}},
			want: 2,
		},
		{args: &TreeNode{Val: 1}, want: 1},
		{args: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, want: 2},
	}

	for i, testCase := range cases {
		t.Run(fmt.Sprintf("TestMinDepth-%d", i), func(t *testing.T) {
			if get := minDepth(testCase.args); testCase.want != get {
				t.Errorf("args: %v, want %v, get %v", inorderTraversal(testCase.args),
					testCase.want, get)
			}
		})
	}
}
