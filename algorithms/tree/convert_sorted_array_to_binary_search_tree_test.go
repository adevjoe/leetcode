package tree

import (
	"fmt"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	testCases := []struct {
		args []int
		want *TreeNode
		get  *TreeNode
	}{
		{
			args: []int{-10, -3, 0, 5, 9},
			want: &TreeNode{Val: 0,
				Left:  &TreeNode{Val: -3, Left: &TreeNode{Val: -10}},
				Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}}},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestSortedArrayToBST-%d", i), func(t *testing.T) {
			if get := sortedArrayToBST(testCase.args); !isSameTree(testCase.want, get) {
				t.Errorf("args: %v, want %v, get %v", testCase.args,
					inorderTraversal(testCase.want), inorderTraversal(get))
			}
		})
	}
}
