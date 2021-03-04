package leetcode

import (
	"reflect"
	"testing"
)

func TestBinaryTreeRightSideView(t *testing.T) {
	cases := []struct {
		desc string
		arg  *TreeNode
		want []int
	}{
		{
			desc: "#1",
			arg: &TreeNode{
				Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}},
			},
			want: []int{1, 3, 4},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := rightSideView(s.arg); !reflect.DeepEqual(got, s.want) {
				t.Errorf("want: %v, got: %v", s.want, got)
			}
		})
	}
}
