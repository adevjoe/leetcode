// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if (p.Val <= root.Val && root.Val <= q.Val) || (p.Val >= root.Val && root.Val >= q.Val) {
		return root
	}
	if root.Left != nil {
		l := lowestCommonAncestor(root.Left, p, q)
		if l != nil {
			return l
		}
	}
	if root.Right != nil {
		r := lowestCommonAncestor(root.Right, p, q)
		if r != nil {
			return r
		}
	}
	return nil
}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if (p.Val <= root.Val && root.Val <= q.Val) || (p.Val >= root.Val && root.Val >= q.Val) {
		return root
	}
	if root.Val > p.Val && root.Val > q.Val {
		if root.Left != nil {
			l := lowestCommonAncestor(root.Left, p, q)
			if l != nil {
				return l
			}
		}
	}
	if root.Val < p.Val && root.Val < q.Val {
		if root.Right != nil {
			r := lowestCommonAncestor(root.Right, p, q)
			if r != nil {
				return r
			}
		}
	}
	return nil
}
