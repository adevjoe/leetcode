package algorithms

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	l := Constructor(2)
	l.Put(1, 1)
	printDoubleList(l.head, l.tail)
	l.Put(2, 2)
	printDoubleList(l.head, l.tail)
	if got := l.Get(1); got != 1 {
		t.Errorf("set 1 error, got: %d", got)
	}
	l.Put(3, 3)
	printDoubleList(l.head, l.tail)
	if got := l.Get(2); got != -1 {
		t.Errorf("evict 1 error, got: %d", got)
	}
	l.Put(4, 4)
	if got := l.Get(1); got != -1 {
		t.Errorf("evict 1 error, got: %d", got)
	}
	if got := l.Get(3); got != 3 {
		t.Errorf("set 3 error, got: %d", got)
	}
	if got := l.Get(4); got != 4 {
		t.Errorf("set 4 error, got: %d", got)
	}
}

func printDoubleList(head, tail *lruDoubleListNode) {
	s := ""
	n := ""
	for head != nil {
		if s != "" {
			s += "->"
		}
		s = fmt.Sprintf("%s%d", s, head.value)
		head = head.next
	}
	for tail != nil {
		if n != "" {
			n += "->"
		}
		n = fmt.Sprintf("%s%d", n, tail.value)
		tail = tail.pre
	}
	fmt.Println(s, "|", n)
}
