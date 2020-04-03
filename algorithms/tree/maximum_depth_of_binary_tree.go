// https://leetcode.com/problems/maximum-depth-of-binary-tree/

package tree

import "math"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}
