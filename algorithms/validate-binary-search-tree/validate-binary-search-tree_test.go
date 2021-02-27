package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestValidateBinarySearchTree(t *testing.T) {
	cases := []struct {
		desc string
		root *TreeNode
		want bool
	}{
		{
			desc: "#1",
			root: &TreeNode{
				Val:  5,
				Left: &TreeNode{Val: 1},
				Right: &TreeNode{
					Val:   4,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 6},
				},
			},
			want: false,
		},
		{
			desc: "#2",
			root: &TreeNode{Val: 1},
			want: true,
		},
		{
			desc: "empty",
			root: &TreeNode{},
			want: true,
		},
		{
			desc: "nil",
			root: nil,
			want: true,
		},
		{
			desc: "#4",
			root: &TreeNode{Val: 2,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: 1},
			},
			want: false,
		},
		{
			desc: "#5",
			root: &TreeNode{Val: 1,
				Left: &TreeNode{Val: 1},
			},
			want: false,
		},
		{
			desc: "#6",
			root: &TreeNode{Val: 0,
				Left: &TreeNode{Val: -1},
			},
			want: true,
		},
	}

	for i, node := range generateTrees(3) {
		cases = append(cases, struct {
			desc string
			root *TreeNode
			want bool
		}{desc: fmt.Sprintf("%d-%d", 3, i), root: node, want: true})
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := isValidBST(s.root); !reflect.DeepEqual(got, s.want) {
				t.Errorf("error")
			}
		})
	}
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return gen(1, n)
}

func gen(start, end int) []*TreeNode {
	var trees []*TreeNode
	if start > end {
		trees = append(trees, nil)
		return trees
	}

	if start == end {
		trees = append(trees, &TreeNode{
			Val:   start,
			Left:  nil,
			Right: nil,
		})
		return trees
	}

	for i := start; i <= end; i++ {
		left := gen(start, i-1)
		right := gen(i+1, end)

		for _, lNode := range left {
			for _, rNode := range right {
				root := &TreeNode{
					Val:   i,
					Left:  lNode,
					Right: rNode,
				}
				trees = append(trees, root)
			}
		}
	}
	return trees
}
