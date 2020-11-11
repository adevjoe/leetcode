// https://leetcode.com/problems/kth-smallest-element-in-a-bst/

package tree

import "math"

type inc struct {
	val int
}

func kthSmallest(root *TreeNode, k int) int {
	inc := &inc{1}
	return kthSmallest1(root, k, inc)
}

func kthSmallest1(root *TreeNode, k int, inc *inc) int {
	if root.Left != nil {
		l := kthSmallest1(root.Left, k, inc)
		if l > math.MinInt64 {
			return l
		}
	}
	if k == inc.val {
		return root.Val
	}
	inc.val++
	if root.Right != nil {
		r := kthSmallest1(root.Right, k, inc)
		if r > math.MinInt64 {
			return r
		}
	}
	return math.MinInt64
}
