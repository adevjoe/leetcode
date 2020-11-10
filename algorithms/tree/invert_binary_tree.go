// https://leetcode.com/problems/invert-binary-tree/

package tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	l := root.Left
	r := root.Right
	root.Right = invertTree(l)
	root.Left = invertTree(r)
	return root
}
