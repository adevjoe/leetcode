// https://leetcode.com/problems/binary-tree-right-side-view/

package tree

import "github.com/adevjoe/leetcode/common"

// solution: level order, get last value
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := common.NewQueue()
	q.Put(root)
	var result []int
	for q.Len() > 0 {
		l := q.Len()
		for i := 1; i <= l; i++ {
			a := q.Pop().(*TreeNode)
			if a.Left != nil {
				q.Put(a.Left)
			}
			if a.Right != nil {
				q.Put(a.Right)
			}
			if i == l {
				result = append(result, a.Val)
			}
		}
	}
	return result
}
