package rotate_string

func leftRotateStringIterator(s string, n int) string {
	if n == 0 {
		return s
	}
	return leftRotateStringIterator(s[1:]+string(s[0]), n-1)
}

// 三步反转法
// 1. abcdef 分为 ab 和 cdef
// 2. 分别反转得到 ba 和 fedc
// 3. 结合 ba 和 fedc 得到 bafedc，然后再反转得到 cdefab
func leftRotateStringByStepThree(s string, n int) string {
	if n >= len(s) {
		return reverseString(s)
	}
	s1, s2 := reverseString(s[:n]), reverseString(s[n:])
	return reverseString(s1 + s2)
}

func reverseString(s string) string {
	b := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
	return string(b)
}
