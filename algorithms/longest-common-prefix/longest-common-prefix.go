// https://leetcode.com/problems/longest-common-prefix

package leetcode

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
