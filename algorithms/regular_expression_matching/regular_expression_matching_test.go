package regular_expression_matching

import (
	"fmt"
	"testing"
)

func TestIsMatch(t *testing.T) {
	fmt.Println("false ->", IsMatch("aa", "a"))
	fmt.Println("true ->", IsMatch("aa", "a*"))
	fmt.Println("true ->", IsMatch("ab", ".*"))
	fmt.Println("true ->", IsMatch("aab", "c*a*b"))
	fmt.Println("false ->", IsMatch("mississippi", "mis*is*p*."))
}
