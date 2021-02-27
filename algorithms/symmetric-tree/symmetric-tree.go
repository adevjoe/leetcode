// https://leetcode.com/problems/symmetric-tree

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var queue []*TreeNode
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		var temp []*int
		l := len(queue)
		for i := 0; i < l; i++ {
			if queue[i] == nil {
				temp = append(temp, nil)
				continue
			} else {
				temp = append(temp, &queue[i].Val)
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			} else {
				queue = append(queue, nil)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			} else {
				queue = append(queue, nil)
			}
		}
		queue = queue[l:]
		if !isSymmetricSlice(temp) {
			return false
		}
	}
	return true
}

func isSymmetricSlice(s []*int) bool {
	if len(s)%2 != 0 {
		return false
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] == nil {
			if s[len(s)-i-1] != nil {
				return false
			}
			continue
		}
		if s[len(s)-i-1] == nil {
			return false
		}
		if *s[i] != *s[len(s)-i-1] {
			return false
		}
	}
	return true
}
