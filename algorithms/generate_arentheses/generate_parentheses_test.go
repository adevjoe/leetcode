package generate_arentheses

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	reflect.DeepEqual([]string{"()"}, generateParenthesis(1))
	reflect.DeepEqual([]string{"(())", "()()"}, generateParenthesis(2))
	reflect.DeepEqual([]string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(2))

	fmt.Println(generateParenthesis2(1))
	fmt.Println(generateParenthesis2(2))
	fmt.Println(generateParenthesis2(3))
}
