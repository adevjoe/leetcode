// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 效率低
func lowestCommonAncestorBT(root, p, q *TreeNode) *TreeNode {
	l := hasNode(root.Left, p, q)
	r := hasNode(root.Right, p, q)
	if l == 1 && r == 1 {
		return root
	}
	if l == 2 {
		return lowestCommonAncestorBT(root.Left, p, q)
	}
	if r == 2 {
		return lowestCommonAncestorBT(root.Right, p, q)
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	return nil
}

func hasNode(root, p, q *TreeNode) int {
	if root == nil {
		return 0
	}
	var i int
	if root.Val == p.Val {
		i++
	}
	if root.Val == q.Val {
		i++
	}
	i += hasNode(root.Left, p, q)
	i += hasNode(root.Right, p, q)
	return i
}

// 效率高
func lowestCommonAncestorBT2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == q.Val || root.Val == p.Val {
		return root
	}
	l := lowestCommonAncestor(root.Left, p, q)
	r := lowestCommonAncestor(root.Right, p, q)
	if l == nil {
		return r
	} else {
		if r == nil {
			return l
		} else {
			return root
		}
	}
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
