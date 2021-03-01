package leetcode

import (
	"reflect"
	"testing"
)

func TestBinaryTreeLevelOrderTraversalII(t *testing.T) {
	cases := []struct {
		desc string
		root *TreeNode
		want [][]int
	}{
		{
			desc: "#1",
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			want: [][]int{{15, 7}, {9, 20}, {3}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := levelOrderBottom(s.root); !reflect.DeepEqual(got, s.want) {
				t.Error()
			}
		})
	}
}
