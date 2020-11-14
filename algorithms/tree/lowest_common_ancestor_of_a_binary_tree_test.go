package tree

import "testing"

func TestLowestCommonAncestorBT(t *testing.T) {
	cases := []struct {
		desc string
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		{
			desc: "#1",
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}},
			p:    &TreeNode{Val: 5},
			q:    &TreeNode{Val: 1},
			want: &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}},
		},
		{
			desc: "#2",
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}},
			p:    &TreeNode{Val: 5},
			q:    &TreeNode{Val: 4},
			want: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}},
		},
	}

	for _, s := range cases {
		t.Run(s.desc, func(t *testing.T) {
			if got := lowestCommonAncestorBT(s.root, s.p, s.q); !isSameTree(got, s.want) {
				t.Errorf("want: %v, got: %v", inorderTraversal(s.want), inorderTraversal(got))
			}
			if got := lowestCommonAncestorBT2(s.root, s.p, s.q); !isSameTree(got, s.want) {
				t.Errorf("want: %v, got: %v", inorderTraversal(s.want), inorderTraversal(got))
			}
		})
	}
}
