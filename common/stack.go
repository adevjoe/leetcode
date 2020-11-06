package common

type (
	NodeStack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func NewNodeStack() *NodeStack {
	return &NodeStack{nil, 0}
}

// Return the number of items in the stack
func (s *NodeStack) Len() int {
	return s.length
}

// View the top item on the stack
func (s *NodeStack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Pop the top item of the stack and return it
func (s *NodeStack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push a value onto the top of the stack
func (s *NodeStack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

// --------------------------------------------------------------
// stack with slice

type SliceStack struct {
	data []interface{}
}

func NewSliceStack() *SliceStack {
	return &SliceStack{
		data: []interface{}{},
	}
}

func (s *SliceStack) Push(v interface{}) {
	s.data = append(s.data, nil)
	s.data = append(s.data[1:], s.data[0:s.Len()-1])
	s.data[0] = v
	return
}

func (s *SliceStack) Pop() (v interface{}) {
	if s.Len() == 0 {
		return nil
	}
	v = s.data[0]
	if s.Len() == 1 {
		s.data = nil
	} else {
		s.data = append(s.data[:s.Len()-2], s.data[1:])
	}
	return v
}

func (s *SliceStack) Peek() (v interface{}) {
	if s.Len() == 0 {
		return nil
	}
	v = s.data[0]
	return v
}

func (s *SliceStack) Len() int {
	return len(s.data)
}
