// https://leetcode.com/problems/populating-next-right-pointers-in-each-node

package leetcode

import "github.com/adevjoe/leetcode/common"

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
	q := common.NewQueue()
	i := 0
	q.Put(root)
	for q.Len() > 0 {
		i++
		//fmt.Println(fmt.Sprintf("line %d: ", i))
		for l := q.Len(); l > 0; l-- {
			v := q.Pop().(*Node)
			// get next
			if l == 1 {
				v.Next = nil
			} else {
				v.Next = q.Touch().(*Node)
			}
			//fmt.Println(v.Val)
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
