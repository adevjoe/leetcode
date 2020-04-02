// https://leetcode.com/problems/validate-binary-search-tree/

package tree

import "math"

func isValidBST(root *TreeNode) bool {
	preVal := new(int)
	first := new(bool)
	*first = true
	return validTraversal(root, preVal, first)
}

func validTraversal(root *TreeNode, preVal *int, first *bool) bool {
	if root == nil {
		return true
	}
	if root.Left != nil {
		if !validTraversal(root.Left, preVal, first) {
			return false
		}
	}
	if root.Val <= *preVal && !*first {
		return false
	}
	*preVal = root.Val
	*first = false
	if root.Right != nil {
		if !validTraversal(root.Right, preVal, first) {
			return false
		}
	}
	return true
}

func isValidBST2(root *TreeNode) bool {
	return recur(root, math.MinInt64, math.MaxInt64)
}

func recur(root *TreeNode, low, high int) bool {
	if root == nil {
		return true
	}

	if root.Val <= low || root.Val >= high {
		return false
	}

	if root.Left != nil {
		if !recur(root.Left, low, root.Val) {
			return false
		}
	}
	if root.Right != nil {
		if !recur(root.Right, root.Val, high) {
			return false
		}
	}
	return true
}
