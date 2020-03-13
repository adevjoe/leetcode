package tree

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(preorderTraversal(root))
	fmt.Println(inorderTraversal(root))
	fmt.Println(postorderTraversal(root))
	fmt.Println(levelOrder(root))
	fmt.Println(inorderTraversal(nil))
}
