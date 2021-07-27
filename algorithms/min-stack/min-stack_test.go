package leetcode

import "testing"

func TestMinStack(t *testing.T) {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	if got := stack.GetMin(); got != -3 {
		t.Error()
	}
	stack.Pop()
	if got := stack.Top(); got != 0 {
		t.Error()
	}
	if got := stack.GetMin(); got != -2 {
		t.Error()
	}
}
