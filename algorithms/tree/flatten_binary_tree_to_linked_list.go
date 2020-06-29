// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

package tree

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	if root.Left != nil {
		if root.Right != nil {
			insert(root.Left, root.Right)
		}
		root.Right = root.Left
		root.Left = nil
	}
}

func insert(root *TreeNode, right *TreeNode) {
	if root.Right != nil {
		insert(root.Right, right)
	} else {
		root.Right = right
	}
}
