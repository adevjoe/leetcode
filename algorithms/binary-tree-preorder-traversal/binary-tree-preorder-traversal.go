// https://leetcode.com/problems/binary-tree-preorder-traversal

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	if root.Left != nil {
		result = append(result, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		result = append(result, preorderTraversal(root.Right)...)
	}
	return result
}
