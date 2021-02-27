// https://leetcode.com/problems/binary-tree-level-order-traversal

package leetcode

import "github.com/adevjoe/leetcode/common"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

func levelOrderWithQueue(root *TreeNode) [][]int {
	s := common.NewQueue()
	s.Put(root)
	var result [][]int
	for s.Len() > 0 {
		var temp []int
		l := s.Len()
		for i := 1; i <= l; i++ {
			t, _ := s.Pop().(*TreeNode)
			if t == nil {
				continue
			}
			if t.Left != nil {
				s.Put(t.Left)
			}
			if t.Right != nil {
				s.Put(t.Right)
			}
			temp = append(temp, t.Val)
		}
		if len(temp) > 0 {
			result = append(result, temp)
		}
	}
	return result
}
