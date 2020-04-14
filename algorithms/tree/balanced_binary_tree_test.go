package tree

import (
	"strconv"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	testCases := []struct {
		args *TreeNode
		want bool
	}{
		{
			args: &TreeNode{Val: 3, Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			want: true,
		},
		{
			args: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
				Right: &TreeNode{Val: 2},
			},
			want: false,
		},
		{
			args: &TreeNode{Val: 1,
				Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}}},
				Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
			},
			want: false,
		},
	}

	for i, testCase := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if get := isBalanced(testCase.args); get != testCase.want {
				t.Errorf("args: %v, get: %v, want: %v", testCase.args,
					get, testCase.want)
			}
		})
	}
}
