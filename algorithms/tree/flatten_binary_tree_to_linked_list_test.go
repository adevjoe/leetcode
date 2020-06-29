package tree

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T) {
	cases := []struct {
		args *TreeNode
		want *TreeNode
	}{
		{
			args: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
				Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}},
			},
			want: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}}}},
		},
	}

	for i, testCase := range cases {
		t.Run(fmt.Sprintf("TestFlatten-%d", i), func(t *testing.T) {
			if flatten(testCase.args); !isSameTree(testCase.args, testCase.want) {
				t.Errorf("want %v, get %v", inorderTraversal(testCase.want), inorderTraversal(testCase.args))
			}
		})
	}
}
