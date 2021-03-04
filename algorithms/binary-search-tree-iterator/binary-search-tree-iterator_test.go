package leetcode

import (
	"fmt"
	"testing"
)

func TestBinarySearchTreeIterator(t *testing.T) {
	tree1 := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	case1 := Constructor(tree1)
	fmt.Println(case1.Next())
	fmt.Println(case1.HasNext())
	fmt.Println(case1.Next())
	fmt.Println(case1.HasNext())
	fmt.Println(case1.Next())
	fmt.Println(case1.HasNext())
	fmt.Println(case1.Next())
	fmt.Println(case1.HasNext())
	fmt.Println(case1.Next())
	fmt.Println(case1.HasNext())
}
