// https://leetcode.com/problems/regular-expression-matching/

package regular_expression_matching

func isMatch(s string, p string) bool {
	dot, asterisk := byte(46), byte(42)
	if p == "" {
		return s == ""
	}
	firstMatch := s != "" && (p[0] == s[0] || p[0] == dot)

	if len(p) >= 2 && p[1] == asterisk {
		return (firstMatch && isMatch(s[1:], p)) || isMatch(s, p[2:])
	} else {
		return firstMatch && isMatch(s[1:], p[1:])
	}
}
