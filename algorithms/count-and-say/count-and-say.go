// https://leetcode.com/problems/count-and-say

package leetcode

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return say(countAndSay(n - 1))
}

func say(s string) string {
	last := uint8(0)
	count := 0
	result := strings.Builder{}
	for i := range s {
		if s[i] == last {
			count++
			continue
		} else {
			if i == 0 {
				last = s[i]
				count = 1
				continue
			} else {
				result.WriteString(strconv.Itoa(count))
				result.WriteString(string(last))
				last = s[i]
				count = 1
			}
		}
	}
	if count > 0 {
		result.WriteString(strconv.Itoa(count))
		result.WriteString(string(last))
	}
	return result.String()
}
