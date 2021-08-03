// https://leetcode.com/problems/factorial-trailing-zeroes

package leetcode

func trailingZeroes(n int) int {
	count := 0
	a := n / 5
	for a > 0 {
		count += a
		a /= 5
	}
	return count
}
