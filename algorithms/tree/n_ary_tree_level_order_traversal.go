// https://leetcode.com/problems/n-ary-tree-level-order-traversal/

package tree

import "github.com/adevjoe/leetcode/common"

type NodeOne struct {
	Val      int
	Children []*NodeOne
}

func levelOrder(root *NodeOne) [][]int {
	if root == nil {
		return nil
	}
	queue := common.NewQueue()
	queue.Put(root)
	var nums [][]int
	for !queue.Empty() {
		var tempNums []int
		l := queue.Len()
		for i := 1; i <= l; i++ {
			r := queue.Pop().(*NodeOne)
			tempNums = append(tempNums, r.Val)
			for _, child := range r.Children {
				queue.Put(child)
			}
		}
		nums = append(nums, tempNums)
	}
	return nums
}
