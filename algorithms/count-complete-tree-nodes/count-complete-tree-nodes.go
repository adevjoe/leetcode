// https://leetcode.com/problems/count-complete-tree-nodes

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	var n int
	if root == nil {
		return 0
	}
	n = 1
	if root.Left != nil {
		n += countNodes(root.Left)
	}
	if root.Right != nil {
		n += countNodes(root.Right)
	}
	return n
}
