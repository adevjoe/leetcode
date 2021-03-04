// https://leetcode.com/problems/sum-root-to-leaf-numbers

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return sumNumbersWithCur(root, 0, 0)
}

func sumNumbersWithCur(root *TreeNode, cur int, sum int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		sum += cur*10 + root.Val
	}
	if root.Left != nil {
		sum = sumNumbersWithCur(root.Left, cur*10+root.Val, sum)
	}
	if root.Right != nil {
		sum = sumNumbersWithCur(root.Right, cur*10+root.Val, sum)
	}
	return sum
}
