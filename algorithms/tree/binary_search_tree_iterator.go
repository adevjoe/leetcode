// https://leetcode.com/problems/binary-search-tree-iterator/

package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	data []int
	cur  int
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		data: inorderTraversal(root),
	}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	if this.HasNext() {
		next := this.data[this.cur]
		this.cur += 1
		return next
	}
	return 0
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.cur <= len(this.data)-1 && this.data != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
