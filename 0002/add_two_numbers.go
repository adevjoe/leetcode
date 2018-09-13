/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a := l1.Val + l2.Val
	if a >= 10 {
		l1.Val = a - 10
		if l2.Next == nil {
			l2.Next = new(ListNode)
		}
		// 有隐患， l2 的值会大于9
		l2.Next.Val += 1
	} else {
		l1.Val = a
	}
	if l2.Next != nil {
		if l1.Next == nil {
			l1.Next = new(ListNode)
		}
		l1.Next = addTwoNumbers(l1.Next, l2.Next)
	}
	return l1
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	printListNode(l1)
	printListNode(l2)
	result := addTwoNumbers(l1, l2)
	printListNode(result)
}

func printListNode(node *ListNode) {
	fmt.Print(node.Val)
	if node.Next != nil {
		fmt.Print(" -> ")
		printListNode(node.Next)
	} else {
		fmt.Print("\n")
	}
}
