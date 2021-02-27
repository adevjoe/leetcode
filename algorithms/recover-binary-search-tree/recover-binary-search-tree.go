// https://leetcode.com/problems/recover-binary-search-tree

package leetcode

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 这题关键是只有2个元素位置错乱，使用中序遍历就可以找出这两个元素，
// 然后把两个元素的值互换。使用递归进行遍历在时间上复杂度是 O(n)，
// 但由于递归使用了栈，在空间上的复杂度为 O(n)，不满足题目的要求。
// 看了网友的的解决思路，可以使用 morris 遍历实现，这样在时间是 O(n)，
// 在空间上则是 O(1)。https://www.cnblogs.com/AnnieKim/archive/2013/06/15/morristraversal.html

// 使用 morris 进行中序遍历，空间复杂度 O(1)
func recoverTree(root *TreeNode) {
	var (
		cur    = root
		pre    *TreeNode
		last   = &TreeNode{Val: math.MinInt64}
		first  *TreeNode
		second *TreeNode
	)

	for cur != nil {
		if cur.Left == nil {
			// get val
			if first == nil && last.Val >= cur.Val {
				first = last
			}
			if first != nil && last.Val >= cur.Val {
				second = cur
			}
			last = cur
			cur = cur.Right
		} else {
			// find predecessor
			pre = cur.Left
			for pre.Right != nil && pre.Right != cur {
				pre = pre.Right
			}

			if pre.Right == nil {
				pre.Right = cur
				cur = cur.Left
			} else {
				pre.Right = nil
				// get val
				if first == nil && last.Val >= cur.Val {
					first = last
				}
				if first != nil && last.Val >= cur.Val {
					second = cur
				}
				last = cur
				cur = cur.Right
			}
		}
	}
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

var (
	first      *TreeNode
	second     *TreeNode
	preElement = &TreeNode{Val: math.MinInt64}
)

// 递归，空间复杂度 O(n)
// for in order traversal
func recoverTree1(root *TreeNode) {
	first, second = nil, nil
	preElement = &TreeNode{Val: math.MinInt64}
	// in order traversal to find the two element
	recoverTreeInOrderTraverse(root)

	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}

func recoverTreeInOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}

	recoverTreeInOrderTraverse(root.Left)

	if first == nil && preElement.Val >= root.Val {
		first = preElement
	}
	if first != nil && preElement.Val >= root.Val {
		second = root
	}
	preElement = root

	recoverTreeInOrderTraverse(root.Right)
}
