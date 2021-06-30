// https://leetcode.com/problems/minimum-window-substring

package leetcode

import "math"

func minWindow(s string, t string) string {
	var (
		start    int
		end      int
		count    int
		minStart int
		minlen   = math.MaxInt64
	)
	cache := make(map[uint8]int)
	for i := range t {
		cache[t[i]]++
	}
	for end < len(s) {
		if cache[s[end]] > 0 {
			count++
		}
		cache[s[end]]--
		for count == len(t) {
			if end-start < minlen {
				minStart = start
				minlen = end - start
			}
			cache[s[start]]++
			if cache[s[start]] > 0 {
				count--
			}
			start++
		}
		end++
	}
	if minlen == math.MaxInt64 {
		return ""
	}
	return s[minStart : minStart+minlen+1]
}
