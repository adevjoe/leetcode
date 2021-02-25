// https://leetcode.com/problems/implement-strstr/

package leetcode

// What is strStr? http://www.cplusplus.com/reference/cstring/strstr/
// Such as Java's indexOf or Go's strings.Index()
func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}
	p := 0
	for i := 0; i < len(haystack); {
		if len(haystack[i:]) < len(needle[p:]) {
			return -1
		}
		if haystack[i] == needle[p] {
			if p == len(needle)-1 {
				return i - len(needle) + 1
			}
			i++
			p++
		} else {
			i = i - p + 1
			p = 0
		}
	}
	return -1
}
