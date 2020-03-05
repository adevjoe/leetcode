// https://leetcode.com/problems/valid-parentheses/

package valid_parentheses

import "github.com/adevjoe/leetcode/common"

const (
	LeftBracket        = "("
	RightBracket       = ")"
	LeftBigBracket     = "{"
	RightBigBracket    = "}"
	LeftMiddleBracket  = "["
	RightMiddleBracket = "]"
)

func isValid(s string) bool {
	stack := common.NewStack()
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
