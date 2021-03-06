package leetcode

import (
	"reflect"
	"testing"
)

func TestInvertBinaryTree(t *testing.T) {
	cases := []struct {
		desc string
		arg  *TreeNode
		want *TreeNode
	}{
		{
			desc: "#1",
			arg:  &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}},
			want: &TreeNode{Val: 4, Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := invertTree(s.arg); !reflect.DeepEqual(got, s.want) {
				t.Error()
			}
		})
	}
}
