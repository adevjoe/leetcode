/*
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.
Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
Example 4:

Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".
Example 5:

Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
*/
package main

import "fmt"

func main() {
	fmt.Println("false ->", isMatch("aa", "a"))
	fmt.Println("true ->", isMatch("aa", "a*"))
	fmt.Println("true ->", isMatch("ab", ".*"))
	fmt.Println("true ->", isMatch("aab", "c*a*b"))
	fmt.Println("false ->", isMatch("mississippi", "mis*is*p*."))
}

func isMatch(s string, p string) bool {
	dot, asterisk := byte(46), byte(42)
	if p == "" {
		return s == ""
	}
	firstMatch := s != "" && (p[0] == s[0] || p[0] == dot)

	if len(p) >= 2 && p[1] == asterisk {
		return (firstMatch && isMatch(s[1:], p)) || isMatch(s, p[2:])
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}
