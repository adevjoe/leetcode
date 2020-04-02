// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

package tree

import "github.com/adevjoe/leetcode/common"

func zigzagLevelOrder(root *TreeNode) [][]int {
	s := common.NewQueue()
	s.Put(root)
	var result [][]int
	flag := false
	for s.Len() > 0 {
		var temp []int
		l := s.Len()
		for i := 1; i <= l; i++ {
			t, _ := s.Pop().(*TreeNode)
			if t == nil {
				continue
			}
			if t.Left != nil {
				s.Put(t.Left)
			}
			if t.Right != nil {
				s.Put(t.Right)
			}
			if flag {
				temp = append(temp, 0)
				copy(temp[1:], temp[0:])
				temp[0] = t.Val
			} else {
				temp = append(temp, t.Val)
			}
		}
		flag = !flag
		if len(temp) > 0 {
			result = append(result, temp)
		}
	}
	return result
}
