package leetcode

import (
	"reflect"
	"testing"
)

func TestConvertSortedArraytoBinarySearchTree(t *testing.T) {
	cases := []struct {
		desc string
		args []int
		want *TreeNode
	}{
		{
			desc: "#1",
			args: []int{-10, -3, 0, 5, 9},
			want: &TreeNode{Val: 0,
				Left:  &TreeNode{Val: -3, Left: &TreeNode{Val: -10}},
				Right: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := sortedArrayToBST(s.args); !reflect.DeepEqual(got, s.want) {
				t.Error()
			}
		})
	}
}
