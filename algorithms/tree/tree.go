package tree

func preorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	result = append(result, root.Val)
	if root.Left != nil {
		result = append(result, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		result = append(result, preorderTraversal(root.Right)...)
	}
	return result
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	if root.Left != nil {
		result = append(result, inorderTraversal(root.Left)...)
	}
	result = append(result, root.Val)
	if root.Right != nil {
		result = append(result, inorderTraversal(root.Right)...)
	}
	return result
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	if root.Left != nil {
		result = append(result, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		result = append(result, postorderTraversal(root.Right)...)
	}
	result = append(result, root.Val)
	return result
}

func levelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	var result [][]int
	queue = append(queue, root)
	for len(queue) > 0 {
		var temp []int
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i] == nil {
				continue
			}
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
		if len(temp) > 0 {
			result = append(result, temp)
		}
	}
	return result
}
