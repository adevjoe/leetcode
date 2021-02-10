// https://leetcode.com/problems/merge-k-sorted-lists/

package merge_k_sorted_lists

import (
	"sort"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	n := []int{}
	for _, l := range lists {
		for l != nil {
			n = append(n, l.Val)
			l = l.Next
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	var head *ListNode
	for _, v := range n {
		tmp := head
		head = &ListNode{
			Val:  v,
			Next: tmp,
		}
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
