// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	s := newQueue()
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

type Queue struct {
	values []interface{}
}

func newQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Len() int {
	return len(q.values)
}

func (q *Queue) Empty() bool {
	return q.Len() == 0
}

func (q *Queue) Put(i interface{}) {
	q.values = append(q.values, i)
}

func (q *Queue) Pop() interface{} {
	if q.Empty() {
		return nil
	}
	v := q.values[0]
	q.values = q.values[1:]
	return v
}
