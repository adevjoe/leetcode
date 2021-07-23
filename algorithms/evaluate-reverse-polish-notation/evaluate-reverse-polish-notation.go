// https://leetcode.com/problems/evaluate-reverse-polish-notation

package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	var s []int
	for i := range tokens {
		if tokens[i] == plus {
			v := s[len(s)-2] + s[len(s)-1]
			s = append(s[:len(s)-2], v)
			continue
		}
		if tokens[i] == sub {
			v := s[len(s)-2] - s[len(s)-1]
			s = append(s[:len(s)-2], v)
			continue
		}
		if tokens[i] == multi {
			v := s[len(s)-2] * s[len(s)-1]
			s = append(s[:len(s)-2], v)
			continue
		}
		if tokens[i] == divide {
			v := s[len(s)-2] / s[len(s)-1]
			s = append(s[:len(s)-2], v)
			continue
		}
		v, _ := strconv.Atoi(tokens[i])
		s = append(s, v)
	}
	return s[0]
}

const (
	plus   = "+"
	sub    = "-"
	multi  = "*"
	divide = "/"
)
