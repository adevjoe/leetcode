// https://leetcode.com/problems/word-break

package leetcode

var m map[string]bool
var ss map[int]bool

func wordBreak(s string, wordDict []string) bool {
	m = make(map[string]bool, 0)
	ss = make(map[int]bool, 0)
	for i := range wordDict {
		m[wordDict[i]] = true
	}
	ss[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < len(s); j++ {
			if ss[j] && m[s[j:i]] {
				ss[i] = true
				break
			}
		}
	}
	return ss[len(s)]
}
