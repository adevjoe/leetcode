// https://leetcode.com/problems/balanced-binary-tree

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := treeHeight(root.Left)
	rightHeight := treeHeight(root.Right)
	if leftHeight-rightHeight > 1 || leftHeight-rightHeight < -1 {
		return false
	}
	if !isBalanced(root.Left) || !isBalanced(root.Right) {
		return false
	}
	return true
}

func treeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	h := 1
	l := treeHeight(root.Left)
	r := treeHeight(root.Right)
	if l > r {
		h += l
	} else {
		h += r
	}
	return h
}
