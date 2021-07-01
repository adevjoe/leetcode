// https://leetcode.com/problems/valid-palindrome

package leetcode

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if !valid(s[i]) {
			i++
			continue
		}
		if !valid(s[j]) {
			j--
			continue
		}
		if lower(s[i]) != lower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func valid(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}
	if b >= 'A' && b <= 'Z' {
		return true
	}
	if b >= '0' && b <= '9' {
		return true
	}
	return false
}

func lower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b + 32
	}
	return b
}
