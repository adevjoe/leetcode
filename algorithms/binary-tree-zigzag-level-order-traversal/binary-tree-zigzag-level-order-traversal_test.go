package leetcode

import (
	"reflect"
	"testing"
)

func TestBinaryTreeZigzagLevelOrderTraversal(t *testing.T) {
	cases := []struct {
		desc string
		root *TreeNode
		want [][]int
	}{
		{
			desc: "#1",
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20,
				Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			want: [][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := zigzagLevelOrder(s.root); !reflect.DeepEqual(got, s.want) {
				t.Error()
			}
		})
	}
}
