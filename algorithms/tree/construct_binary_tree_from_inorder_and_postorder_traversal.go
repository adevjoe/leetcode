// https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/

package tree

func buildTreePostOrder(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	i := index(inorder, postorder[len(postorder)-1])
	if i == -1 {
		return nil
	}
	root := &TreeNode{Val: inorder[i]}
	if i > 0 {
		// left
		root.Left = buildTreePostOrder(inorder[:i], findSame(inorder[:i], postorder[:len(postorder)-1]))
	}
	if i+1 < len(inorder) {
		// right
		root.Right = buildTreePostOrder(inorder[i+1:], findRight(inorder[i+1:], postorder[:len(postorder)-1]))
	}
	return root
}

func findSame(inOrder, postOrder []int) []int {
	for pi, p := range postOrder {
		if index(inOrder, p) == -1 {
			return postOrder[:pi]
		}
		if pi == len(postOrder)-1 {
			return postOrder
		}
	}
	return nil
}

func findRight(inOrder, postOrder []int) []int {
	for pi, p := range postOrder {
		if index(inOrder, p) != -1 {
			return postOrder[pi:]
		}
	}
	return nil
}

func buildTreePostOrder2(in []int, po []int) *TreeNode {
	return dfs(0, len(in)-1, len(in), in, po)
}

func dfs(inp, pop, l int, in []int, po []int) *TreeNode {
	if l <= 0 {
		return nil
	}
	root := TreeNode{Val: po[pop]}
	if l == 1 {
		return &root
	}

	root_in_in := inp
	for in[root_in_in] != po[pop] {
		root_in_in += 1
	}
	left_len := root_in_in - inp
	right_len := l - left_len - 1
	root.Left = dfs(inp, pop-right_len-1, left_len, in, po)
	root.Right = dfs(root_in_in+1, pop-1, right_len, in, po)

	return &root
}
