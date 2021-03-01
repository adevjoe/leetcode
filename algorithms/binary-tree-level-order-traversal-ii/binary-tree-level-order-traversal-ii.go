// https://leetcode.com/problems/binary-tree-level-order-traversal-ii

package leetcode

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	s := NewQueue()
	s.Put(root)
	var result [][]int
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
			temp = append(temp, t.Val)
		}
		if len(temp) > 0 {
			result = append(result, []int{})
			copy(result[1:], result[0:])
			result[0] = temp
		}
	}
	return result
}

type Queue struct {
	values []interface{}
}

func NewQueue() *Queue {
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

func (q *Queue) GetIndex(i int) interface{} {
	if q.Len() <= i {
		return nil
	}
	v := q.values[i]
	q.values = append(q.values[:i], q.values[i+1:]...)
	return v
}
