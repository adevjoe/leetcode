/**
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:
Input: ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Note:
All given inputs are in lowercase letters a-z.
*/
package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	i := 0
	over := false
	ss := ""
	l := len(strs)
	if l == 0 {
		return ""
	}
	for !over {
		temp := uint8(0)
		for key, s := range strs {
			if len(s)-1 < i {
				over = true
				break
			}
			if temp == 0 {
				temp = s[i]
			} else if temp != s[i] {
				over = true
				break
			}
			if key == l-1 {
				ss += string(temp)
				i++
			}
		}
	}
	return ss
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{}))
}
