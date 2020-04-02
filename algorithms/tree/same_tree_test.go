package tree

import "testing"

func TestSameTree(t *testing.T) {
	t1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	t2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	if !isSameTree(t1, t2) {
		t.Error()
	}
	t3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	t4 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	if isSameTree(t3, t4) {
		t.Error()
	}
	t5 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}
	t6 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	if isSameTree(t5, t6) {
		t.Error()
	}
}
