// https://leetcode.com/problems/minimum-depth-of-binary-tree/

package tree

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minLeft := minDepth(root.Left)
	minRight := minDepth(root.Right)
	if minLeft == 0 {
		return 1 + minRight
	}
	if minRight == 0 {
		return 1 + minLeft
	}
	return 1 + int(math.Min(float64(minLeft), float64(minRight)))
}
