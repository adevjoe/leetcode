// https://leetcode.com/problems/binary-tree-maximum-path-sum

package leetcode

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Input: [-10,9,20,null,null,15,7]

   -10
   / \
  9  20
    /  \
   15    7

Output: 42
*/

func maxPathSum(root *TreeNode) int {
	_, max := maxPathDown(root, math.MinInt64)
	return max
}

func maxPathDown(root *TreeNode, max int) (int, int) {
	if root == nil {
		return 0, max
	}
	l, max := maxPathDown(root.Left, max)
	left := math.Max(0, float64(l))
	r, max := maxPathDown(root.Right, max)
	right := math.Max(0, float64(r))
	max = int(math.Max(float64(max), left+right+float64(root.Val)))
	return int(math.Max(left, right)) + root.Val, max
}
