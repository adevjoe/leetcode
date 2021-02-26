package leetcode

import (
	"fmt"
	"testing"
	"time"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	s := "abba"
	s1 := time.Now().UnixNano()
	fmt.Println(lengthOfLongestSubstring(s))
	s2 := time.Now().UnixNano()
	fmt.Println(s2 - s1)
	fmt.Println(lengthOfLongestSubstring2(s))
	s3 := time.Now().UnixNano()
	fmt.Println(s3 - s2)
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	s := "pwwkew"
	for n := 0; n < b.N; n++ {
		lengthOfLongestSubstring(s)
	}
}

func BenchmarkLengthOfLongestSubstring2(b *testing.B) {
	s := "pwwkew"
	for n := 0; n < b.N; n++ {
		lengthOfLongestSubstring2(s)
	}
}
