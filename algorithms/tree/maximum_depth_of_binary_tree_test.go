package tree

import "testing"

func TestMaxDepth(t *testing.T) {
	t1 := &TreeNode{Val: 3, Left: &TreeNode{Val: 9},
		Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	if get := maxDepth(t1); get != 3 {
		t.Errorf("want 3, get %d", get)
	}
	var t2 *TreeNode
	if maxDepth(t2) != 0 {
		t.Error()
	}
}
