// https://leetcode.com/problems/add-strings/submissions/
// TODO Not Good, think a better solution
package string_add

import "strings"

func stringAdd(num1, num2 string) string {
	if !validStringNumber(num1) || !validStringNumber(num2) {
		return ""
	}
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	lena := len(num1)
	lenb := len(num2)
	s := ""
	cur := lena - 1
	up := false
	for cur >= 0 {
		var num uint8
		if ai := cur - (lena - lenb); ai >= 0 {
			num = num1[cur] - '0' + num2[ai]
		} else {
			num = num1[cur]
		}
		if up {
			num += 1
			up = false
		}
		if num > '9' {
			up = true
			s = string(num-10) + s
		} else {
			s = string(num) + s
		}
		cur--
	}
	if up {
		s = "1" + s
	}
	return s
}

func stringAddForByte(num1, num2 string) string {
	if !validStringNumber(num1) || !validStringNumber(num2) {
		return ""
	}
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	lena := len(num1)
	lenb := len(num2)
	var s []byte
	cur := lena - 1
	up := false
	for cur >= 0 {
		var num uint8
		if ai := cur - (lena - lenb); ai >= 0 {
			num = num1[cur] - '0' + num2[ai]
		} else {
			num = num1[cur]
		}
		if up {
			num += 1
			up = false
		}
		if num > '9' {
			up = true
			s = append([]byte{num - 10}, s...)
		} else {
			s = append([]byte{num}, s...)
		}
		cur--
	}
	if up {
		s = append([]byte{'1'}, s...)
	}
	sb := strings.Builder{}
	sb.Write(s)
	return sb.String()
}

func validStringNumber(s string) bool {
	for _, ss := range s {
		if ss < '0' || ss > '9' {
			return false
		}
	}
	if len(s) >= 2 && s[0] == '0' {
		return false
	}
	return true
}
