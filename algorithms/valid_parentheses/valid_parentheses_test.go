package valid_parentheses

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
