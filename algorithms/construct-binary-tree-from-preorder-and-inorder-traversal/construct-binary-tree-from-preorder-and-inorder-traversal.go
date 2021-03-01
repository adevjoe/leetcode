// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeWithPoint(&preorder, inorder)
}

func buildTreeWithPoint(preorder *[]int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	p := *preorder
	index := index(inorder, p[0])
	*preorder = p[1:]
	root := &TreeNode{Val: inorder[index]}
	root.Left = buildTreeWithPoint(preorder, inorder[:index])
	root.Right = buildTreeWithPoint(preorder, inorder[index+1:])
	return root
}

func index(list []int, value int) int {
	for i, v := range list {
		if v == value {
			return i
		}
	}
	return -1
}
