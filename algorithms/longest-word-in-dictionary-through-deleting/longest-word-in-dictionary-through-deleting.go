// https://leetcode.com/problems/longest-word-in-dictionary-through-deleting

package leetcode

func findLongestWord(s string, dictionary []string) string {
	longest := ""
	for _, s2 := range dictionary {
		if len(s2) > len(longest) {
			if containWord(s, s2) {
				longest = s2
			}
		}
		if len(s2) == len(longest) {
			if !containWord(s, s2) {
				continue
			}
			if s2 < longest {
				longest = s2
			}
		}
	}
	return longest
}

func containWord(s1, s2 string) bool {
	if len(s1) < len(s2) {
		return false
	}
	if len(s1) == len(s2) && s1 == s2 {
		return true
	}
	var cur1, cur2 int
	for cur2 < len(s2) && cur1 < len(s1) {
		if s1[cur1] == s2[cur2] {
			if cur2 == len(s2)-1 {
				return true
			}
			cur1++
			cur2++
			continue
		}
		cur1++
		if len(s1[cur1:]) < len(s2[cur2:]) {
			return false
		}
	}
	return false
}
