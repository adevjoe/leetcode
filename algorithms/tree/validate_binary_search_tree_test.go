package tree

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	for _, node := range generateTrees(3) {
		fmt.Println(isValidBST(node) == true)
	}
	fmt.Println(isValidBST(&TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 6},
		},
	}) == false)
	fmt.Println(isValidBST(&TreeNode{Val: 1}) == true)
	fmt.Println(isValidBST(&TreeNode{}) == true)
	fmt.Println(isValidBST(nil) == true)
	fmt.Println(isValidBST(&TreeNode{Val: 2,
		Left:  &TreeNode{Val: 3},
		Right: &TreeNode{Val: 1},
	}) == false)
	fmt.Println(isValidBST(&TreeNode{Val: 1,
		Left: &TreeNode{Val: 1},
	}) == false)
	fmt.Println(isValidBST(&TreeNode{Val: 0,
		Left: &TreeNode{Val: -1},
	}) == true)
}
