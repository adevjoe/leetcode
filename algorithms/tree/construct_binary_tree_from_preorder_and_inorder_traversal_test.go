package tree

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preorder, inorder)
	fmt.Println(preorder)
	fmt.Println(preorderTraversal(tree))
	fmt.Println(inorder)
	fmt.Println(inorderTraversal(tree))
}
