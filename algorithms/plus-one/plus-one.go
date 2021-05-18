// https://leetcode.com/problems/plus-one

package leetcode

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	l := len(digits) + 1
	result := make([]int, l, l)
	result[0] = 1
	return result
}
