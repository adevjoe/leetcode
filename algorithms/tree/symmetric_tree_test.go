package tree

import (
	"testing"
)

func TestIsSymmetricTree(t *testing.T) {
	t1 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 3}}}
	t2 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Right: &TreeNode{Val: 3}},
		Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
	t3 := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3}}
	if !isSymmetric(t1) {
		t.Error()
	}
	if isSymmetric(t2) {
		t.Error()
	}
	if isSymmetric(t3) {
		t.Error()
	}
}
