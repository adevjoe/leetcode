package leetcode

import (
	"fmt"
	"testing"
)

func TestRegularExpressionMatching(t *testing.T) {
	fmt.Println("false ->", isMatch("aa", "a"))
	fmt.Println("true ->", isMatch("aa", "a*"))
	fmt.Println("true ->", isMatch("ab", ".*"))
	fmt.Println("true ->", isMatch("aab", "c*a*b"))
	fmt.Println("false ->", isMatch("mississippi", "mis*is*p*."))
}
