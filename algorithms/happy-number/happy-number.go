// https://leetcode.com/problems/happy-number

package leetcode

func isHappy(n int) bool {
	result := 0
	for n != 1 {
		if n < 6 {
			return false
		}
		for n > 0 {
			tmp := n % 10
			result += tmp * tmp
			n /= 10
		}
		n = result
		if n == 1 {
			return true
		}
		result = 0
	}
	return true
}

func isHappy2(n int) bool {
	sum := 0
	for n >= 10 {
		sum = 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
	}
	return n == 1 || n == 7
}
