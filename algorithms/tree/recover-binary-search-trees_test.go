package tree

import "testing"

func TestRecoverBST(t *testing.T) {
	t1 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}},
	}
	t.Logf("old t1: %v", inorderTraversal(t1))
	recoverTree(t1)
	t.Logf("new t1: %v", inorderTraversal(t1))

	t2 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}},
	}
	t.Logf("old t1: %v", inorderTraversal(t2))
	recoverTree(t2)
	t.Logf("new t1: %v", inorderTraversal(t2))
	t3 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}},
	}
	recoverTree(t3)
}

func TestRecoverBST1(t *testing.T) {
	t1 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}},
	}
	t.Logf("old t1: %v", inorderTraversal(t1))
	recoverTree1(t1)
	t.Logf("new t1: %v", inorderTraversal(t1))

	t2 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}},
	}
	t.Logf("old t1: %v", inorderTraversal(t2))
	recoverTree1(t2)
	t.Logf("new t1: %v", inorderTraversal(t2))
}
