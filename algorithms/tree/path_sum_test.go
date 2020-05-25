package tree

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {
	cases := []struct {
		arg1 *TreeNode
		arg2 int
		want bool
	}{
		{
			arg1: &TreeNode{Val: 5,
				Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
				Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}},
			},
			arg2: 22,
			want: true,
		},
		{
			arg1: &TreeNode{Val: 5,
				Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
				Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}},
			},
			arg2: 32,
			want: false,
		},
		{
			arg1: &TreeNode{Val: 5,
				Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
				Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}},
			},
			arg2: 26,
			want: true,
		},
		{
			arg1: &TreeNode{Val: 1},
			arg2: 1,
			want: true,
		},
	}

	for i, testCase := range cases {
		t.Run(fmt.Sprintf("TestPathSum-%d", i), func(t *testing.T) {
			if get := hasPathSum(testCase.arg1, testCase.arg2); testCase.want != get {
				t.Errorf("arg1: %v, arg2: %v, want %v, get %v", inorderTraversal(testCase.arg1),
					testCase.arg2, testCase.want, get)
			}
		})
	}
}
