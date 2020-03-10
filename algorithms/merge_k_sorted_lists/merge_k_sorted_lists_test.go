package merge_k_sorted_lists

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	var list []*ListNode
	list = append(list, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	})
	list = append(list, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	})
	list = append(list, &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	})
	for _, node := range list {
		printListNode(node)
	}
	printListNode(mergeKLists(list))
}

func printListNode(head *ListNode) {
	fmt.Print(head.Val)
	if head.Next != nil {
		fmt.Print(" -> ")
		printListNode(head.Next)
	} else {
		fmt.Print("\n")
	}
}
