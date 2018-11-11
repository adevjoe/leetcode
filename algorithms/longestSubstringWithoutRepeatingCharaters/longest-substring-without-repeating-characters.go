/**
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", which the length is 3.

Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
package longestSubstringWithoutRepeatingCharaters

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	longestSubString := ""
	currentString := ""
	hash := map[int]int{}
	for key, value := range s {
		if key == 0 {
			hash = map[int]int{int(value): 1}
			currentString = string(value)
			longestSubString = currentString
			continue
		}
		if hash[int(value)] > 0 { // repeated
			currentString = s[hash[int(value)] : key+1]
			start := hash[int(value)]
			hash = map[int]int{}
			for i := start; i < key+1; i++ {
				hash[int(s[i])] = i + 1
			}
		} else {
			hash[int(value)] = key + 1
			currentString += string(value)
		}
		if len(currentString) > len(longestSubString) {
			longestSubString = currentString
		}
		fmt.Println(currentString, longestSubString, hash)
	}
	return len(longestSubString)
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 1
	cur := 1
	start := 0

	for i := 1; i < len(s); i++ {
		c := s[i]

		// check repeated char
		found := false
		for j := i - 1; j >= start; j-- {
			if s[j] == c {
				cur = i - j
				start = j + 1
				found = true
				break
			}
		}

		// step
		if !found {
			cur++
		}
		if cur > max {
			max = cur
		}
	}

	return max
}

func main() {
	s := "abba"
	s1 := time.Now().UnixNano()
	fmt.Println(lengthOfLongestSubstring(s))
	s2 := time.Now().UnixNano()
	fmt.Println(s2 - s1)
	fmt.Println(lengthOfLongestSubstring2(s))
	s3 := time.Now().UnixNano()
	fmt.Println(s3 - s2)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	for true {
		lengthOfLongestSubstring(s)
	}
	select {}
}
