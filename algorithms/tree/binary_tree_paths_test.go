package tree

import (
	"reflect"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {
	cases := []struct {
		desc string
		arg  *TreeNode
		want []string
	}{
		{
			desc: "#1",
			arg: &TreeNode{
				Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3},
			},
			want: []string{"1->2->5", "1->3"},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := binaryTreePaths(s.arg); !reflect.DeepEqual(got, s.want) {
				t.Errorf("arg: %v, want: %v, got: %v", inorderTraversal(s.arg), s.want, got)
			}
		})
	}
}
