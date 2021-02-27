package leetcode

import (
	"fmt"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
