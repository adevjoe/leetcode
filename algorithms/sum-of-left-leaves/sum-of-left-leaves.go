// https://leetcode.com/problems/sum-of-left-leaves

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return sumOfLeftLeavesSub(root, false)
}

func sumOfLeftLeavesSub(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}
	if isLeft && root.Left == nil && root.Right == nil {
		return root.Val
	}

	return sumOfLeftLeavesSub(root.Left, true) + sumOfLeftLeavesSub(root.Right, false)
}
