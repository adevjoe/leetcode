// https://leetcode.com/problems/duplicate-zeros/

package leetcode

func duplicateZeros(arr []int) {
	s := newQueue()
	for i := 0; i < len(arr); i++ {
		cur := arr[i]
		if s.len() > 0 {
			r := s.pop()
			arr[i] = r
			s.push(cur)
		}
		if cur == 0 {
			s.push(0)
		}
	}
}

type queue struct {
	vals []int
}

func newQueue() *queue {
	return &queue{}
}

func (s *queue) push(v int) {
	s.vals = append(s.vals, v)
}

func (s *queue) pop() int {
	val := s.vals[0]
	s.vals = s.vals[1:]
	return val
}

func (s *queue) len() int {
	return len(s.vals)
}

func duplicateZeros2(arr []int) {
	i := 0
	sh := 0
	for i = 0; sh+i < len(arr); i++ {
		if arr[i] == 0 {
			sh++
		}
	}
	for i := i - 1; sh > 0; i-- {
		if i+sh < len(arr) {
			arr[i+sh] = arr[i]
		}
		if arr[i] == 0 {
			sh--
			arr[i+sh] = arr[i]
		}
	}
}
