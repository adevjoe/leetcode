package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructBinaryTreefromPreorderandInorderTraversal(t *testing.T) {
	cases := []struct {
		desc     string
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{
			desc:     "#1",
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want:     &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
		},
		{
			desc:     "#2",
			preorder: []int{-1},
			inorder:  []int{-1},
			want:     &TreeNode{Val: -1},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := buildTree(s.preorder, s.inorder); !reflect.DeepEqual(got, s.want) {
				t.Error()
			}
		})
	}
}
