package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructBinaryTreefromInorderandPostorderTraversal(t *testing.T) {
	cases := []struct {
		desc      string
		inorder   []int
		postorder []int
		want      *TreeNode
	}{
		{
			desc:    "#1",
			inorder: []int{9, 3, 15, 20, 7}, postorder: []int{9, 15, 7, 20, 3},
			want: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
		},
		{
			desc:    "#2",
			inorder: []int{2, 3, 1}, postorder: []int{3, 2, 1},
			want: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
		},
		{
			desc:    "#3",
			inorder: []int{3, 2, 1}, postorder: []int{3, 2, 1},
			want: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
		},
		{
			desc:    "#4",
			inorder: []int{1, 2, 3}, postorder: []int{3, 2, 1},
			want: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := buildTreePostOrder(s.inorder, s.postorder); !reflect.DeepEqual(got, s.want) {
				t.Error()
			}
			if got := buildTreePostOrder2(s.inorder, s.postorder); !reflect.DeepEqual(got, s.want) {
				t.Error()
			}
		})
	}
}
