// https://leetcode.com/problems/valid-parentheses/

package valid_parentheses

const (
	LeftBracket        = "("
	RightBracket       = ")"
	LeftBigBracket     = "{"
	RightBigBracket    = "}"
	LeftMiddleBracket  = "["
	RightMiddleBracket = "]"
)

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (s *Stack) Len() int {
	return s.length
}

// View the top item on the stack
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Pop the top item of the stack and return it
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push a value onto the top of the stack
func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

func isValid(s string) bool {
	stack := New()
	for _, v := range s {
		if v == '{' {
			stack.Push('}')
		} else if v == '[' {
			stack.Push(']')
		} else if v == '(' {
			stack.Push(')')
		} else if stack.Len() == 0 || stack.Pop() != v {
			return false
		}
	}
	return stack.Len() == 0
}
