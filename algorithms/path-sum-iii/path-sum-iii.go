// https://leetcode.com/problems/path-sum-iii

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	return pathSumSub(root, sum) + pathSumWarp(root, sum, true)
}

func pathSumWarp(root *TreeNode, sum int, isRoot bool) int {
	if root == nil {
		return 0
	}
	var n int
	if !isRoot {
		n += pathSumSub(root, sum)
	}
	return pathSumWarp(root.Left, sum, false) + pathSumWarp(root.Right, sum, false) + n

}

func pathSumSub(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var n int
	n += pathSumSub(root.Left, sum-root.Val)
	n += pathSumSub(root.Right, sum-root.Val)
	if root.Val == sum {
		n += 1
	}
	return n
}
