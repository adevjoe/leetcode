// https://leetcode.com/problems/wildcard-matching

package leetcode

func isMatch(s string, p string) bool {
	pi := 0
	ss := -1
	ps := -1
	for i := 0; i < len(s); i++ {
		if pi > len(p)-1 {
			if ss >= 0 && ps >= 0 {
				pi = ps
				i = ss
				continue
			}
			return false
		}
		if p[pi] == '*' {
			ss = i
			ps = pi
			i--
			pi++
		} else {
			if s[i] == p[pi] || p[pi] == '?' {
				pi++
				continue
			} else {
				if ss >= 0 && ps >= 0 {
					pi = ps
					i = ss
				} else {
					return false
				}
			}
		}
	}
	for pi < len(p) && p[pi] == '*' {
		pi++
	}
	return pi == len(p)
}
