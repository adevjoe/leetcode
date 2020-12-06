// https://leetcode.com/problems/path-sum-iii/

package tree

func pathSumIII(root *TreeNode, sum int) int {
	return pathSumIIISub(root, sum) + pathSumWarp(root, sum, true)
}

func pathSumWarp(root *TreeNode, sum int, isRoot bool) int {
	if root == nil {
		return 0
	}
	var n int
	if !isRoot {
		n += pathSumIIISub(root, sum)
	}
	return pathSumWarp(root.Left, sum, false) + pathSumWarp(root.Right, sum, false) + n

}

func pathSumIIISub(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	var n int
	n += pathSumIIISub(root.Left, sum-root.Val)
	n += pathSumIIISub(root.Right, sum-root.Val)
	if root.Val == sum {
		n += 1
	}
	return n
}
