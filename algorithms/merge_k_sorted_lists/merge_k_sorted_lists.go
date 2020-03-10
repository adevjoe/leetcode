// https://leetcode.com/problems/merge-k-sorted-lists/

package merge_k_sorted_lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 先连接所有链表，然后排序
func mergeKLists(lists []*ListNode) *ListNode {
	var result *ListNode
	var tail *ListNode
	for _, list := range lists {
		if result == nil {
			result = list
		} else {
			tail.Next = list
		}
		for list != nil {
			tail = list
			list = list.Next
		}
	}
	return result
}

type ListNode struct {
	Val  int
	Next *ListNode
}
