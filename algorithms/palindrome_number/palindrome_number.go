// https://leetcode.com/problems/palindrome-number/

package palindrome_number

func IsPalindrome(x int) bool {
	// filter
	// -121 is false
	if x < 0 {
		return false
	}
	// 10 110 4440 ... false
	if x%10 == 0 && x != 0 {
		return false
	}

	new := 0
	current := x

	// convert position backward as forward
	for new < x {
		new = new*10 + current%10
		current /= 10
	}

	if x != new {
		return false
	}
	return true
}
