package tree

import (
	"fmt"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	t1 := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20,
		Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	/**
	[
	  [3],
	  [20,9],
	  [15,7]
	]
	*/
	fmt.Println(zigzagLevelOrder(t1))
}
