// https://leetcode.com/problems/copy-list-with-random-pointer

package leetcode

// Node Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// solution 1
/**
origin node: 7 -> 13(7) -> 11(1) -> 10(11) -> 1(7)
loop1: 7 -> 7 -> 13(7) -> 13 -> 11(1) -> 11 -> 10(11) -> 10 -> 1(7) -> 1
loop2: 7 -> 7 -> 13(7) -> 13(7) -> 11(1) -> 11(1) -> 10(11) -> 10(11) -> 1(7) -> 1
loop3: 7 -> 13(7) -> 11(1) -> 10(11) -> 1(7)
*/
func copyRandomList(head *Node) *Node {
	var result *Node

	// loop 1, duplicate each node
	loop := head
	for loop != nil {
		next := loop.Next
		loop.Next = &Node{
			Val:  loop.Val,
			Next: next,
		}
		loop = loop.Next.Next
	}

	// add random for duplicate node
	loop = head
	for loop != nil {
		if loop.Random != nil {
			loop.Next.Random = loop.Random.Next
		}
		loop = loop.Next.Next
	}

	// pick up duplicate node
	loop = head
	if head != nil {
		result = head.Next
	}
	for loop != nil {
		if loop.Next.Next == nil {
			loop.Next = nil
			break
		}
		tmp := loop.Next
		loop.Next = loop.Next.Next
		loop = tmp
	}
	return result
}

// use map
/**
loop1: put all node to map, origin node point as key, new node point as value
loop2: fill all new node's next node and random node
*/
func copyRandomList2(head *Node) *Node {
	m := make(map[*Node]*Node)
	loop := head
	for loop != nil {
		m[loop] = &Node{
			Val: loop.Val,
		}
		loop = loop.Next
	}
	loop = head
	for loop != nil {
		if loop.Next != nil {
			m[loop].Next = m[loop.Next]
		}
		if loop.Random != nil {
			m[loop].Random = m[loop.Random]
		}
		loop = loop.Next
	}
	return m[head]
}
