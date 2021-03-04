// https://leetcode.com/problems/house-robber-iii

package leetcode

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var robMap map[*TreeNode]map[bool]int

func rob(root *TreeNode) int {
	return robSub(root, true)
}

func robTwo(root *TreeNode) int {
	robMap = make(map[*TreeNode]map[bool]int)
	return robTwoSub(root, true)
}

func robSub(root *TreeNode, haveRoot bool) int {
	if root == nil {
		return 0
	}
	if haveRoot {
		a := root.Val + robSub(root.Left, false) + robSub(root.Right, false)
		b := robSub(root.Left, true) + robSub(root.Right, true)
		return int(math.Max(float64(a), float64(b)))
	} else {
		return robSub(root.Left, true) + robSub(root.Right, true)
	}
}

func robTwoSub(root *TreeNode, haveRoot bool) int {
	if root == nil {
		return 0
	}
	if robMap[root][haveRoot] > 0 {
		return robMap[root][haveRoot]
	}
	var max int
	if haveRoot {
		a := root.Val + robTwoSub(root.Left, false) + robTwoSub(root.Right, false)
		b := robTwoSub(root.Left, true) + robTwoSub(root.Right, true)
		max = int(math.Max(float64(a), float64(b)))
	} else {
		max = robTwoSub(root.Left, true) + robTwoSub(root.Right, true)
	}
	if robMap[root] == nil {
		robMap[root] = make(map[bool]int)
	}
	robMap[root][haveRoot] = max
	return max
}
