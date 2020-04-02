package tree

import (
	"fmt"
	"testing"
)

func TestGenerateTrees(t *testing.T) {
	for _, node := range generateTrees(1) {
		fmt.Println(preorderTraversal(node))
	}

	fmt.Println("------------------")

	for _, node := range generateTrees(2) {
		fmt.Println(preorderTraversal(node))
	}

	fmt.Println("------------------")

	for _, node := range generateTrees(3) {
		fmt.Println(preorderTraversal(node))
	}
}
