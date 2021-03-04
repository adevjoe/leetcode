// https://leetcode.com/problems/add-strings

package leetcode

import "strings"

func addStrings(num1 string, num2 string) string {
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
