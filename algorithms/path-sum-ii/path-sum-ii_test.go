package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPathSumII(t *testing.T) {
	cases := []struct {
		arg1 *TreeNode
		arg2 int
		want [][]int
	}{
		{
			arg1: &TreeNode{Val: 5,
				Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
				Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}},
			},
			arg2: 22,
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
		{
			arg1: &TreeNode{Val: 1},
			arg2: 1,
			want: [][]int{{1}},
		},
	}

	for i, testCase := range cases {
		t.Run(fmt.Sprintf("TestPathSum-%d", i), func(t *testing.T) {
			if get := pathSum(testCase.arg1, testCase.arg2); !reflect.DeepEqual(testCase.want, get) {
				t.Error()
			}
		})
	}
}
