// https://leetcode.com/problems/populating-next-right-pointers-in-each-node-ii

package leetcode

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	q := NewQueue()
	i := 0
	q.Put(root)
	for q.Len() > 0 {
		i++
		for l := q.Len(); l > 0; l-- {
			v := q.Pop().(*Node)
			// get next
			if l == 1 {
				v.Next = nil
			} else {
				v.Next = q.Touch().(*Node)
			}
			if v.Left != nil {
				q.Put(v.Left)
			}
			if v.Right != nil {
				q.Put(v.Right)
			}
		}
	}
	return root
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

func (q *Queue) Touch() interface{} {
	if q.Empty() {
		return nil
	}
	return q.values[0]
}
