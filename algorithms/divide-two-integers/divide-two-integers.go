// https://leetcode.com/problems/divide-two-integers

package leetcode

func divide(dividend int, divisor int) int {
	flag := false
	if dividend < 0 {
		dividend = -dividend
		flag = true
	}
	if divisor < 0 {
		divisor = -divisor
		flag = !flag
	}
	t := 0
	for dividend >= divisor {
		count := 1
		tmp := divisor
		for tmp <= dividend {
			tmp <<= 1
			count <<= 1
		}
		t += count >> 1
		dividend -= tmp >> 1
	}
	if flag {
		t = -t
	}
	if t > 1<<31-1 {
		t = 1<<31 - 1
	}
	if t < -1<<31 {
		t = -1 << 31
	}
	return t
}
