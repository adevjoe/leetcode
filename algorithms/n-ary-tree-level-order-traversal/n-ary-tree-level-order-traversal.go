// https://leetcode.com/problems/n-ary-tree-level-order-traversal

package leetcode

// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	queue := NewQueue()
	queue.Put(root)
	var nums [][]int
	for !queue.Empty() {
		var tempNums []int
		l := queue.Len()
		for i := 1; i <= l; i++ {
			r := queue.Pop().(*Node)
			tempNums = append(tempNums, r.Val)
			for _, child := range r.Children {
				queue.Put(child)
			}
		}
		nums = append(nums, tempNums)
	}
	return nums
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
