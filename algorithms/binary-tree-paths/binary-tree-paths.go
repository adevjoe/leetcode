// https://leetcode.com/problems/binary-tree-paths

package leetcode

import "strconv"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	return binaryTreePathsSub(root, "")
}

func binaryTreePathsSub(root *TreeNode, path string) []string {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{path + strconv.Itoa(root.Val)}
	}
	path += strconv.Itoa(root.Val) + "->"
	l := binaryTreePathsSub(root.Left, path)
	r := binaryTreePathsSub(root.Right, path)
	var result []string
	result = append(result, l...)
	result = append(result, r...)
	return result
}
