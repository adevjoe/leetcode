// https://leetcode.com/problems/invert-binary-tree

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	l := root.Left
	r := root.Right
	root.Right = invertTree(l)
	root.Left = invertTree(r)
	return root
}
