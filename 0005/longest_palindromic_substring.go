/**
Given a string s, find the longest palindromic substring in s. You may assume that
the maximum length of s is 1000.

https://leetcode.com/problems/longest-palindromic-substring/description/
 */
package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}

func longestPalindrome(s string) string {
	l := ""
	for a := 0; a < len(s); a++ {
		for b := a+1; b <= len(s); b++ {
			ss := s[a:b]
			if len(ss) > len(l) && isPalindrome(ss) {
				l = ss
			}
		}
	}
	return l
}

func isPalindrome(s string) bool {
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
