package tree

import (
	"fmt"
	"testing"
)

func TestNumTrees(t *testing.T) {
	fmt.Println(numTrees(0) == 1)
	fmt.Println(numTrees(1) == 1)
	fmt.Println(numTrees(2) == 2)
	fmt.Println(numTrees(3) == 5)

	fmt.Println(numTrees3(0) == 1)
	fmt.Println(numTrees3(1) == 1)
	fmt.Println(numTrees3(2) == 2)
	fmt.Println(numTrees3(3) == 5)
}
