package leetcode

import (
	"fmt"
	"testing"
)

func TestSumRoottoLeafNumbers(t *testing.T) {
	cases := []struct {
		arg1 *TreeNode
		want int
	}{
		{arg1: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, want: 25},
		{arg1: &TreeNode{Val: 4, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 0}}, want: 1026},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("TestSumNumbers-%d", i), func(t *testing.T) {
			if get := sumNumbers(c.arg1); get != c.want {
				t.Error()
			}
		})
	}
}
